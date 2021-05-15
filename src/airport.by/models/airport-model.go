package models

type Airport struct {
  Name string `json:"title"`
}

type Airline struct {
  Name string `json:"title"`
}

type Aircraft struct {
  Name string `json:"title"`
}

type Status struct {
  Code string `json:"id"`
  Name string `json:"title"`
}

type AirportArrival struct {
  Id          string    `json:"flight_id"`
  FlightCode  string    `json:"flight"`
  Airport     Airport   `json:"airport"`
  Airline     Airline   `json:"airline"`
  Aircraft    Aircraft  `json:"aircraft"`
  Status      Status    `json:"status"`
  PlannedTime string    `json:"plan"`
  ActualTime  string    `json:"fact"`
  Gate        string    `json:"gate"`
  IsCancelled bool      `json:"isCancelled"`
  IsDelayed   bool      `json:"isDelayed"`
}
