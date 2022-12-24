// Package entity defines domain/business logic specific models that can be used in any layer.
package entity

type (
	// Weather define the weather entity
	Weather struct {
		Longitude string `json:"longitude"`
		Latitude  string `json:"latitude"`
	}
)

// Weather creates new weather type
func NewWeather() Weather {
	return Weather{
		Longitude: "",
		Latitude:  "",
	}
}
