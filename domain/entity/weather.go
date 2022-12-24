// Package entity defines domain/business logic specific models that can be used in any layer.
package entity

type (
	// Weather define the weather entity
	Weather struct {
		Longitude string `json:"longitude"`
		Latitude  string `json:"latitude"`
	}

	WeatherOutput struct {
		Coord      map[string]float64
		Weather    []InnerWeather
		Base       string
		Main       map[string]float64
		Visibility float64
		Wind       map[string]float64
		Clouds     map[string]float64
		Date       float64
		Sys        Sys
		Timezone   float64
		ID         float64
		Name       string
		Cod        float64
		Alerts     Alerts
	}

	InnerWeather struct {
		ID          float64
		Main        string
		Description string
		Icon        string
	}

	Alerts struct {
		Sender_Name string
		Event       string
		Start       float64
		End         float64
		Description string
	}

	Sys struct {
		Country string
		Sunrise float64
		Sunset  float64
	}
)

// Weather creates new weather type
func NewWeather() Weather {
	return Weather{
		Longitude: "",
		Latitude:  "",
	}
}
