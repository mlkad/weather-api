package models

// Weather represents the weather data structure
type Weather struct {
	ID          int    `json:"id"`
	Temperature int    `json:"temperature"`
	Condition   string `json:"condition"`
}
