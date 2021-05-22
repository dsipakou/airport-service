package main

import (
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
  "encoding/json"
  "os"
  "context"
  "time"

  firebase "firebase.google.com/go"
  //"firebase.google.com/go/auth"

  //"google.golang.org/api/option"
  "github.io/dsipakou/airport-service/pkg/models"
)

func main() {
  arrivalResponse, err := http.Get(os.Getenv("ARRIVAL_URL"))
  departureResponse, depErr := http.Get(os.Getenv("DEPARTURE_URL"))

  if err != nil {
    log.Fatalln(err)
  }

  if depErr != nil {
    log.Fatalln(depErr)
  }
  defer arrivalResponse.Body.Close()
  defer departureResponse.Body.Close()

  var arrivals []models.AirportArrival
  var departures []models.AirportDepartureNow
  var yesterdayArrivals []models.AirportArrivalYesterday
  var todayArrivals []models.AirportArrivalToday
  var tomorrowArrival []models.AirportArrivalTomorrow

  if err := json.NewDecoder(arrivalResponse.Body).Decode(&arrivals); err != nil {
    panic(err)
  }
  if err := json.NewDecoder(departureResponse.Body).Decode(&departures); err != nil {
    panic(err)
  }

  arrivalBody, arrErr := ioutil.ReadAll(arrivalResponse.Body)
  departureBody, depErr := ioutil.ReadAll(departureResponse.Body)

  if arrErr != nil {
    panic(arrErr)
  }

  sb := string(arrivalBody)
  sd := string(departureBody)

  ctx := context.Background()
  config := &firebase.Config{
    DatabaseURL: os.Getenv("DATABASE_URL"),
  }

  app, err := firebase.NewApp(ctx, config)
  if err != nil {
    panic(err)
  }

  client, err := app.Database(ctx)
  if err != nil {
    panic(err)
  }

  // if err := client.NewRef("departures").Set(ctx, departures); err != nil {
  //   panic(err)
  // }

  layout := time.RFC3339
  plan := "2021-05-18T05:45:00+03:00"

  t, err := time.Parse(layout, plan)

  if err != nil {
    panic(err)
  }

  timeNow := time.Now() 

  for _, v := range arrivals {
    localTime, err := time.Parse(layout, v.PlannedTime)
    if err == nil {
      if localTime.Day() == timeNow.Add(-24 * time.Hour).Day() {
        yesterdayArrivals = append(yesterdayArrivals, models.AirportArrivalYesterday{v})
      } else if localTime.Day() == timeNow.Day() {
        todayArrivals = append(todayArrivals, models.AirportArrivalToday{v})
      } else if localTime.Day() == timeNow.Add(24 * time.Hour).Day() {
        tomorrowArrival = append(tomorrowArrival, models.AirportArrivalTomorrow{v})
      }
    } 
    
  }

  if err := client.NewRef("arrivals").Set(ctx, ""); err != nil {
    panic(err)
  }

  if err := client.NewRef("arrivals/yesterday").Set(ctx, yesterdayArrivals); err != nil {
    panic(err)
  }

  if err := client.NewRef("arrivals/today").Set(ctx, todayArrivals); err != nil {
    panic(err)
  }

  if err := client.NewRef("arrivals/tomorrow").Set(ctx, tomorrowArrival); err != nil {
    panic(err)
  }

  fmt.Println(t)
  fmt.Println(timeNow.Add(-96*time.Hour).Day() == t.Day())


  fmt.Println(sb)
  fmt.Println(sd)

  fmt.Println(arrivals[0])
}
