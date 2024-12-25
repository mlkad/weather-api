package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"weather-api/models"
)

// This is a simple model for weather data
var weatherData = []models.Weather{
	{ID: 1, Temperature: 22, Condition: "Sunny"},
	{ID: 2, Temperature: 18, Condition: "Cloudy"},
}

// GetWeather - Handles the GET request for fetching all weather data
func GetWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)
}

// GetWeatherByID - Handles the GET request to fetch weather data by ID
func GetWeatherByID(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL parameters
	id := mux.Vars(r)["id"]

	// Find the weather data with the matching ID
	for _, weather := range weatherData {
		if fmt.Sprintf("%v", weather.ID) == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(weather)
			return
		}
	}

	// If not found, send a 404 response
	http.Error(w, "Weather not found", http.StatusNotFound)
}

// CreateWeather - Handles the POST request to create new weather data
func CreateWeather(w http.ResponseWriter, r *http.Request) {
	var newWeather models.Weather
	// Parse the request body
	if err := json.NewDecoder(r.Body).Decode(&newWeather); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add new weather data to the collection
	weatherData = append(weatherData, newWeather)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newWeather)
}

// UpdateWeather - Handles the PUT request to update existing weather data
func UpdateWeather(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updatedWeather models.Weather

	// Find the weather data with the matching ID
	for i, weather := range weatherData {
		if fmt.Sprintf("%v", weather.ID) == id {
			// Update weather data
			json.NewDecoder(r.Body).Decode(&updatedWeather)
			weatherData[i] = updatedWeather

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedWeather)
			return
		}
	}

	http.Error(w, "Weather not found", http.StatusNotFound)
}

// DeleteWeather - Handles the DELETE request to delete weather data
func DeleteWeather(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	// Find the weather data with the matching ID
	for i, weather := range weatherData {
		if fmt.Sprintf("%v", weather.ID) == id {
			// Remove weather data
			weatherData = append(weatherData[:i], weatherData[i+1:]...)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Weather not found", http.StatusNotFound)
}
