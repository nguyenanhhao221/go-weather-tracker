package main

import (
	"log"
	"net/http"
	"net/url"
)

func queryCityHandler(w http.ResponseWriter, r *http.Request) {
	apiCfg := apiConfig{}
	apiKey := apiCfg.getApiKey()

	city := "London"
	response, err := getCityWeather(apiKey, city)
	if err != nil {
		responseWithJSON(w, response.StatusCode, err)
	}
	defer response.Body.Close()
	responseWithJSON(w, response.StatusCode, response.Body)
}

func getCityWeather(apiKey string, city string) (*http.Response, error) {
	parsedUrl, err := url.Parse(baseUrl)
	if err != nil {
		log.Printf("Failed to parse baseUrl: %v", err)
	}

	params := url.Values{}
	params.Add("appid", apiKey)
	params.Add("q", city)

	parsedUrl.RawQuery = params.Encode()
	log.Println(parsedUrl.String())
	return http.Get(parsedUrl.String())
}
