package routes

import (
	"github.com/gorilla/mux"
	"weather-api/controllers"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Define the routes for the weather API
	r.HandleFunc("/weather", controllers.GetWeather).Methods("GET")            // GET route to fetch weather
	r.HandleFunc("/weather", controllers.CreateWeather).Methods("POST")        // POST route to create new weather data
	r.HandleFunc("/weather/{id}", controllers.GetWeatherByID).Methods("GET")   // GET route to fetch weather by ID
	r.HandleFunc("/weather/{id}", controllers.UpdateWeather).Methods("PUT")    // PUT route to update weather data
	r.HandleFunc("/weather/{id}", controllers.DeleteWeather).Methods("DELETE") // DELETE route to delete weather data

	return r
}
