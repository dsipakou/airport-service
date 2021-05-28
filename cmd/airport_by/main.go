package main

import (
  // "fmt"
  "net/http"
  "log"
  // "io/ioutil"
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
  var departures []models.AirportDeparture
  var yesterdayArrivals []models.AirportArrivalYesterday
  var yesterdayDepartures []models.AirportDepartureYesterday
  var todayArrivals []models.AirportArrivalToday
  var todayDepartures []models.AirportDepartureToday
  var tomorrowArrival []models.AirportArrivalTomorrow
  var tomorrowDeparture []models.AirportDepartureTomorrow

  if err := json.NewDecoder(arrivalResponse.Body).Decode(&arrivals); err != nil {
    panic(err)
  }
  if err := json.NewDecoder(departureResponse.Body).Decode(&departures); err != nil {
    panic(err)
  }

  // arrivalBody, arrErr := ioutil.ReadAll(arrivalResponse.Body)
  // departureBody, depErr := ioutil.ReadAll(departureResponse.Body)

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

  layout := time.RFC3339

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

  for _, v := range departures {
    localTime, err := time.Parse(layout, v.PlannedTime)
    if err == nil {
      if localTime.Day() == timeNow.Add(-24 * time.Hour).Day() {
        yesterdayDepartures = append(yesterdayDepartures, models.AirportDepartureYesterday{v})
      } else if localTime.Day() == timeNow.Day() {
        todayDepartures = append(todayDepartures, models.AirportDepartureToday{v})
      } else if localTime.Day() == timeNow.Add(24 * time.Hour).Day() {
        tomorrowDeparture = append(tomorrowDeparture, models.AirportDepartureTomorrow{v})
      }
    }
  }

  if err := client.NewRef("arrivals").Set(ctx, ""); err != nil {
    panic(err)
  }

  if err := client.NewRef("departures").Set(ctx, ""); err != nil {
    panic(err)
  }

  if err := client.NewRef("arrivals/yesterday").Set(ctx, yesterdayArrivals); err != nil {
    panic(err)
  }

  if err := client.NewRef("departures/yesterday").Set(ctx, yesterdayDepartures); err != nil {
    panic(err)
  }

  if err := client.NewRef("arrivals/today").Set(ctx, todayArrivals); err != nil {
    panic(err)
  }

  if err := client.NewRef("departures/today").Set(ctx, todayDepartures); err != nil {
    panic(err)
  }

  if err := client.NewRef("arrivals/tomorrow").Set(ctx, tomorrowArrival); err != nil {
    panic(err)
  }

  if err := client.NewRef("departures/tomorrow").Set(ctx, tomorrowDeparture); err != nil {
    panic(err)
  }
}
