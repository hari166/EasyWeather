package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	owm "github.com/briandowns/openweathermap"
)

type weatherData struct {
	Temperature float64
	Humidity    int
	Weather     string
	Wind        int
	UVIndex     float64
	Forecast    forecastData
}

type forecastData struct {
	Tempo [5]int `json:"temp"`
}

func getWeather(apikey, location string) (*weatherData, error) {
	w, err := owm.NewCurrent("C", "EN", apikey)
	if err != nil {
		return nil, err
	}
	err = w.CurrentByName(location)
	if err != nil {
		return nil, err
	}
	uvIndexResp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/uvi?appid=%s&lat=%f&lon=%f", apikey, w.GeoPos.Latitude, w.GeoPos.Longitude))
	if err != nil {
		return nil, err
	}
	defer uvIndexResp.Body.Close()
	var uvIndexRespData struct {
		Value float64 `json:"value"`
	}

	if err := json.NewDecoder(uvIndexResp.Body).Decode(&uvIndexRespData); err != nil {
		return nil, err
	}

	//forecast
	lat := w.GeoPos.Latitude
	lon := w.GeoPos.Longitude
	forecast, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s&units=metric", lat, lon, apikey))
	if err != nil {
		return nil, err
	}
	defer forecast.Body.Close()

	var forecastDataSlice [5]int
	type apiResponse struct {
		List []struct {
			Main struct {
				Temp float64 `json:"temp"`
			} `json:"main"`
		} `json:"list"`
	}
	var apiRespData apiResponse
	if err := json.NewDecoder(forecast.Body).Decode(&apiRespData); err != nil {
		return nil, err
	}
	for i := 0; i < 5; i++ {
		forecastDataSlice[i] = int(math.Round(apiRespData.List[i].Main.Temp))
	}
	//end of forecast

	weather := &weatherData{
		Temperature: math.Round(w.Main.Temp),
		Humidity:    w.Main.Humidity,
		Weather:     w.Weather[0].Description,
		Wind:        int(w.Wind.Speed),
		UVIndex:     uvIndexRespData.Value,
		Forecast:    forecastData{Tempo: forecastDataSlice},
	}
	return weather, nil
}
