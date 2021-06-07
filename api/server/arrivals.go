package arrivals

import (
	"net/http"
	"github.io/dsipakou/airport-service/pkg/models"
	firebase "firebase.google.com/go"
	"encoding/json"
	"os"
	"context"
	"time"
)

func readArrivals() {
	arrivalResponse, err := http.Get(os.Getenv("ARRIVAL_URL"))

	if err != nil {
		panic(err)
	}

  	defer arrivalResponse.Body.Close()

	var arrivals []models.AirportArrival
  	var yesterdayArrivals []models.AirportArrivalYesterday
  	var todayArrivals []models.AirportArrivalToday
  	var tomorrowArrival []models.AirportArrivalTomorrow
  	var nowArrivals []models.AirportArrivalNow

	if err := json.NewDecoder(arrivalResponse.Body).Decode(&arrivals); err != nil {
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

			if localTime.Day() == timeNow.Day() && localTime.Hour() == timeNow.Hour() {
				nowArrivals = append(nowArrivals, models.AirportArrivalNow{v})
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

	if err := client.NewRef("arrivals/now").Set(ctx, nowArrivals); err != nil {
		panic(err)
	}
}