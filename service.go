package main

import (
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
)

func main() {
  resp, err := http.Get("http://airport.by/ru/flights/arrival")

  if err != nil {
    log.Fatalln(err)
  }

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    log.Fatalln(err)
  }

  sb := string(body)

  fmt.Println(sb)
}
