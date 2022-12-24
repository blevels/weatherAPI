// Package handler is the API handler for the Get Weather use case
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/blevels/weatherAPI/adapter/api/response"
	"github.com/blevels/weatherAPI/adapter/logger"
	"github.com/blevels/weatherAPI/usecase"
)

type (
	// GetWeatherRequest Request data sent to Open Weather API
	GetWeatherRequest struct {
		Longitude string `json:"longitude"`
		Latitude  string `json:"latitude"`
	}

	// GetWeatherHandler defines the dependencies of the HTTP handler for the use case
	GetWeatherHandler struct {
		uc  usecase.GetWeatherUseCase
		log logger.Logger
	}
)

// NewGetWeatherHandler creates new use case handler with its dependencies
func NewGetWeatherHandler(uc usecase.GetWeatherUseCase, log logger.Logger) GetWeatherHandler {
	return GetWeatherHandler{
		uc:  uc,
		log: log,
	}
}

// Handle handles http request
func (g GetWeatherHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var reqData GetWeatherRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		g.log.WithFields(logger.Fields{
			"error":       err.Error(),
			"http_status": http.StatusBadRequest,
		}).Errorf("failed to marshal message")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	output, err := g.uc.Execute(r.Context(), usecase.GetWeatherInput{
		Longitude: reqData.Longitude,
		Latitude:  reqData.Latitude,
	})
	if err != nil {
		g.log.WithFields(logger.Fields{
			"error":       err.Error(),
			"http_status": http.StatusInternalServerError,
		}).Errorf("error when creating a new transfer")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
}
