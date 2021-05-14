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
  resp, err := http.Get(os.Getenv("ARRIVAL_URL"))

  if err != nil {
    log.Fatalln(err)
  }

  defer resp.Body.Close()

  var arrivals []models.AirportArrival

  if err := json.NewDecoder(resp.Body).Decode(&arrivals); err != nil {
    fmt.Println("Oops")
    panic(err)
  }

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err)
  }

  sb := string(body)

  fmt.Printf("%v", cResp[0])
  fmt.Println(sb)
  fmt.Println(resp)
}
