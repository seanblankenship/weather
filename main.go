package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
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

func getAPIKey() (string, error) {
	godotenv.Load()
	
	key := os.Getenv("WEATHER_API_KEY")
	if key == "" {
		return "", fmt.Errorf("WEATHER_API_KEY environment variable not set. Create a .env file or set the environment variable")
	}
	
	return key, nil
}

func fetchWeather(apiKey, location string) (*Weather, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	baseURL := "https://api.weatherapi.com/v1/forecast.json"
	params := url.Values{}
	params.Add("key", apiKey)
	params.Add("q", location)
	params.Add("aqi", "no")
	params.Add("days", "1")
	params.Add("alerts", "no")
	
	fullURL := baseURL + "?" + params.Encode()
	
	res, err := client.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error making API request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("weather API returned status %d. Check your API key and location", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}
	
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, fmt.Errorf("error parsing weather data: %v", err)
	}
	
	return &weather, nil
}

func displayWeather(weather *Weather) error {
	if len(weather.Forecast.Forecastday) == 0 {
		return fmt.Errorf("no forecast data available")
	}
	
	location, current := weather.Location, weather.Current
	hours := weather.Forecast.Forecastday[0].Hour

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
	
	return nil
}

func main() {
	location := "Tampa"
	if len(os.Args) >= 2 {
		location = os.Args[1]
	}
	
	key, err := getAPIKey()
	if err != nil {
		log.Fatalf("Error getting API key: %v", err)
	}
	
	weather, err := fetchWeather(key, location)
	if err != nil {
		log.Fatalf("%v", err)
	}
	
	if err := displayWeather(weather); err != nil {
		log.Fatalf("Error displaying weather: %v", err)
	}
}
