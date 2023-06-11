package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type apiConfig struct{}

var baseUrl string = "https://api.openweathermap.org/data/2.5/weather"

func (apiCfg *apiConfig) getApiKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error: failed to load env", err)
	}

	openWeatherApiKey := os.Getenv("OPEN_WEATHER_API_KEY")
	return openWeatherApiKey
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error: failed to load env", err)
	}

	PORT := os.Getenv("PORT")

	http.HandleFunc("/query", queryCityHandler)

	log.Println("Server is listening on port:", PORT)
	startServerErr := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal("Error starting server", startServerErr)
	}
}
