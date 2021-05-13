package models

type AirportArrival struct {
  Id string `json:"flight_id"`
  FlightCode string `json:"flight"`
  PlannedTime string `json:"plan"`
  ActualTime string `json:"fact"`
  Gate string `json:"gate"`
  IsCancelled bool `json:"isCancelled"`
  IsDelayed bool `json:"isDelayed"`
}
