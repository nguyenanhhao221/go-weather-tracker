package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	// When take the payload, marshal them from go structure to json
	jsonData, err := json.Marshal(payload)

	if err != nil {
		log.Println("Error: Failed to marshal json payload", err)
		w.WriteHeader(500)
	}

	// Add header to let client know the contain type is json
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	// Write the json data we marshaled to the response
	_, resWriteErr := w.Write(jsonData)
	if resWriteErr != nil {
		log.Println("Error: Failed to write response ", resWriteErr)
		w.WriteHeader(500)
		return
	}
}
func responseWithError(w http.ResponseWriter, statusCode int, msg string) {
	// Handle when back end error
	if statusCode > 499 {
		log.Println("Responding with 5XX error", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	responseWithJSON(w, statusCode, errResponse{
		Error: msg,
	})
}
