// Package http (adapter) provides an HTTP client interface for the entire application
package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/blevels/weatherAPI/adapter/logger"
	"github.com/blevels/weatherAPI/config"
	"github.com/blevels/weatherAPI/domain/entity"
	"github.com/blevels/weatherAPI/usecase"
	"io"
)

type (
	requestSender struct {
		client HttpGetter
		log    logger.Logger
		apiKey string
		uri    string
	}
)

// NewApiRequestSender creates new request sender with its dependencies
func NewApiRequestSender(client HttpGetter, l logger.Logger, cfg config.Weather) usecase.ApiRequestSender {
	return requestSender{
		client: client,
		log:    l,
		apiKey: cfg.Key,
		uri:    cfg.URI,
	}
}

// Send creates the API request to the Open Weather API for the Get Weather use case
func (r requestSender) Send(_ context.Context, w entity.Weather) (map[string]interface{}, error) {
	res, err := r.client.Get(fmt.Sprintf(`%s?lat=%s&lon=%s&appid=%s&units=imperial`, r.uri, w.Latitude, w.Longitude, r.apiKey))
	if err != nil {
		r.log.WithFields(logger.Fields{
			"error": err.Error(),
		}).Errorf("failed to client")
		return map[string]interface{}{}, err
	}

	result := make(map[string]interface{})

	jsonData, err := io.ReadAll(res.Body)
	if err := json.Unmarshal([]byte(jsonData), &result); err != nil {
		return map[string]interface{}{}, err
	}

	r.log.WithFields(logger.Fields{
		"http_status": res.StatusCode,
	}).Infof("success to authorized")

	return result, nil
}
