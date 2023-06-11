package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func queryCityHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		log.Println("cannot find valid city parameter")
		responseWithError(w, 404, "cannot find valid city parameter")
		return
	}

	apiCfg := apiConfig{}
	apiKey := apiCfg.getApiKey()

	log.Printf("Getting weather data for %v", city)
	response, err := getCityWeather(apiKey, city)
	if err != nil {
		log.Println("Error when getting api city", err)
	}
	responseWithJSON(w, 200, response)
}

type CityWeatherRes struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func getCityWeather(apiKey string, city string) (*CityWeatherRes, error) {
	parsedUrl, err := url.Parse(baseUrl)
	if err != nil {
		log.Printf("Failed to parse baseUrl: %v", err)
		return nil, err
	}

	params := url.Values{}
	params.Add("appid", apiKey)
	params.Add("q", city)

	parsedUrl.RawQuery = params.Encode()

	weatherResp, err := http.Get(parsedUrl.String())
	if err != nil {
		return nil, err
	}
	defer weatherResp.Body.Close()

	decoder := json.NewDecoder(weatherResp.Body)
	cityWeatherData := CityWeatherRes{}
	decodeErr := decoder.Decode(&cityWeatherData)

	if decodeErr != nil {
		log.Println("Error: Failed to decode json from OpenWeatherAPI", decodeErr)
		return nil, decodeErr
	}

	return &cityWeatherData, nil
}
