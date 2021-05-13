package main

import (
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
  "encoding/json"
  "airport.by/models"
)

func main() {
  resp, err := http.Get("http://airport.by/en/flights/arrival")

  if err != nil {
    log.Fatalln(err)
  }

  defer resp.Body.Close()


  var cResp []models.AirportArrival

  if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
    fmt.Println("Oops")
    log.Fatalln(err)
  }

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    log.Fatalln(err)
  }

  sb := string(body)

  fmt.Printf("%v", cResp[0])
  fmt.Println(sb)
  fmt.Println(resp)
}
