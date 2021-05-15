package main

import (
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
  "encoding/json"
  "airport.by/models"
  "os"
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

  if err := json.NewDecoder(arrivalResponse.Body).Decode(&arrivals); err != nil {
    fmt.Println("Oops")
    panic(err)
  }
  if err := json.NewDecoder(departureResponse.Body).Decode(&departures); err != nil {
    fmt.Println("Oops")
    panic(err)
  }

  arrivalBody, arrErr := ioutil.ReadAll(arrivalResponse.Body)
  departureBody, depErr := ioutil.ReadAll(departureResponse.Body)

  if arrErr != nil {
    panic(arrErr)
  }

  sb := string(arrivalBody)
  sd := string(departureBody)

  fmt.Printf("%v", arrivals[0])
  fmt.Printf("%v", departures[0])
  fmt.Println(sb)
  fmt.Println(sd)
  fmt.Println(arrivalResponse)
  fmt.Println(departureResponse)
}
