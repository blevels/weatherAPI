package handler

import (
	"github.com/blevels/weatherAPI/adapter/logger"
	"github.com/blevels/weatherAPI/usecase"
	"html/template"
	"net/http"
	"time"
)

type (
	// GetPageHandler defines the dependencies of the HTTP handler for the use case
	GetPageHandler struct {
		log logger.Logger
	}
)

type WeatherInfo struct {
	Headline string
	Body     string
	Success  bool
	Data     usecase.GetWeatherOutput
}

// Template
var pageTmplt *template.Template

// Map name formatDate to formatDate function above
var funcMap = template.FuncMap{
	"formatDate":   formatDate,
	"wDescription": getWeatherDescription,
	"wMain":        getWeatherMain,
}

func getWeatherMain(weather []interface{}) interface{} {
	d, _ := weather[0].(map[string]interface{})
	return d["main"]
}

func getWeatherDescription(weather []interface{}) interface{} {
	d, _ := weather[0].(map[string]interface{})
	return d["description"]
}

// Custom function must have only 1 return value, or 1 return value and an error
func formatDate(timeFloat float64) string {
	//Define layout for formatting timestamp to string
	return time.Unix(int64(timeFloat), int64(0)).Format("Mon, 02 Jan 2006 15:04:05 -0700")
}

// NewGetPageHandler creates new use case handler with its dependencies
func NewGetPageHandler(log logger.Logger) GetPageHandler {
	return GetPageHandler{
		log: log,
	}
}

// Handle handles http request
func (g GetPageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	pageTmplt = template.Must(template.New("form").Funcs(funcMap).ParseFiles("index.html"))
	weather := WeatherInfo{
		Headline: "Weather Service Information",
		Body:     "Please input your coordinates.",
	}

	err := pageTmplt.Execute(w, weather)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Error(w, "", http.StatusBadRequest)
}
