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

type AirportArrivalNow struct {
  AirportArrival
}

type AirportArrivalToday struct {
  AirportArrival
}

type AirportArrivalYesterday struct {
  AirportArrival
}

type AirportArrivalTomorrow struct {
  AirportArrival
}

type AirportDeparture struct {
  Id                  string    `json:"flight_id"`
  FlightCode          string    `json:"flight"`
  Airport             Airport   `json:"airport"`
  Airline             Airline   `json:"airline"`
  PlannedTime         string    `json:"plan"`
  ActualTime          string    `json:"fact"`
  Status              Status    `json:"status"`
  Gate                string    `json:"gate"`
  IsCancelled         bool      `json:"isCancelled"`
  IsDelayed           bool      `json:"isDelayed"`
  DelayedTill         string    `json:"DelayedTo"`
  RegistrationNumbers []string  `json:"numbers_reg"`
  GateNumbers         []string  `json:"numbers_gate"`
}

func (departure AirportDeparture) String() string {
  return departure.Id
}

type AirportDepartureNow struct {
  AirportDeparture
}

type AirportDepartureToday struct {
  AirportDeparture
}

type AirportDepartureYesterday struct {
  AirportDeparture
}

type AirportDepartureTomorrow struct {
  AirportDeparture
}
