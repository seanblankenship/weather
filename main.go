package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"github.com/fatih/color"
)

type Weather struct {
	Location struct {
		Name string `json:"name"`
		Region string `json:"region"`
	} `json:"location"`
	Current struct {
		TempF float64 `json:"temp_f"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64 `json:"time_epoch"`
				TempF float64 `json:"temp_f"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {

	q := "Tampa"
	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	
	key := "***REMOVED***"
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + key + "&q=" + q + "&aqi=no&days=1&alerts=no")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available.")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf(
		"%s, %s: %.0fF, %s\n",
		location.Name,
		location.Region,
		current.TempF,
		current.Condition.Text,
	)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)
		if date.Before(time.Now()) {
			continue
		}
		message := fmt.Sprintf(
			"%s - %.0fF, %.0f%%, %s\n",
			date.Format("3:04 PM"),
			hour.TempF,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)
		if hour.ChanceOfRain < 40 {
			fmt.Print(message)
		} else {
			color.Red(message)
		}

	}
}
