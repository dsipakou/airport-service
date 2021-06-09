package departures

import (
	"net/http"
	"encoding/json"
	"os"
	"context"
	"time"
	firebase "firebase.google.com/go"
	"github.io/dsipakou/airport-service/pkg/models"
)

func ReadDepartures() {
	response, err := http.Get(os.Getenv("DEPARTURE_URL"))

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var departures []models.AirportDeparture
	var yesterdayDepartures []models.AirportDepartureYesterday
	var todayDepartures []models.AirportDepartureToday
	var tomorrowDeparture []models.AirportDepartureTomorrow
	var nowDepartures []models.AirportDepartureNow

	if err := json.NewDecoder(response.Body).Decode(&departures); err != nil {
		panic(err)
	}

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

			if localTime.Day() == timeNow.Day() && localTime.Hour() == timeNow.Hour() {
				nowDepartures = append(nowDepartures, models.AirportDepartureNow{v})
			}
		}
	}

	entryTypeMap := map[string]struct = {"departures/yesterday": yesterdayDepartures}

	if err := client.NewRef("departures").Set(ctx, ""); err != nil {
		panic(err)
	}

	if err := client.NewRef("departures/yesterday").Set(ctx, yesterdayDepartures); err != nil {
		panic(err)
	}

	if err := client.NewRef("departures/today").Set(ctx, todayDepartures); err != nil {
		panic(err)
	}

	if err := client.NewRef("departures/tomorrow").Set(ctx, tomorrowDeparture); err != nil {
		panic(err)
	}

	if err := client.NewRef("departures/now").Set(ctx, nowDepartures); err != nil {
		panic(err)
	}
}