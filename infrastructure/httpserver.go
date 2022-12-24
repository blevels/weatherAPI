// Package infrastruture provides the initialization of the HTTP server, API endpoints and all handlers
package infrastructure

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/blevels/weatherAPI/adapter/api/handler"
	adapterhttp "github.com/blevels/weatherAPI/adapter/http"
	adapterlogger "github.com/blevels/weatherAPI/adapter/logger"
	"github.com/blevels/weatherAPI/adapter/presenter"
	"github.com/blevels/weatherAPI/config"
	infrahttp "github.com/blevels/weatherAPI/infrastructure/http"
	"github.com/blevels/weatherAPI/infrastructure/logger"
	"github.com/blevels/weatherAPI/infrastructure/router"
	"github.com/blevels/weatherAPI/usecase"
)

// HTTPServer define an application structure
type HTTPServer struct {
	logger adapterlogger.Logger
	router router.Router
}

// NewHTTPServer creates new HTTPServer with its dependencies
func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		logger: logger.NewLogrus(),
		router: router.NewMux(),
	}
}

// getWeatherHandler use case handler orchestration
func (h HTTPServer) getWeatherHandler(cfg config.Weather) http.HandlerFunc {
	requestSender := adapterhttp.NewApiRequestSender(
		infrahttp.NewClient(
			infrahttp.NewRequest(
				infrahttp.WithRetry(infrahttp.NewRetry(3, []int{http.StatusInternalServerError}, 400*time.Millisecond)),
				infrahttp.WithTimeout(5*time.Second),
			),
		),
		h.logger,
		cfg,
	)

	uc := usecase.NewGetWeatherInteractor(
		requestSender,
		presenter.NewGetWeatherPresenter(),
	)

	return handler.NewGetWeatherHandler(uc, h.logger).Handle
}

// Start setup routes and run the application server
func (h HTTPServer) Start(cfg *config.Config) {
	h.router.GET("/healthCheck", healthCheck)
	h.router.POST("/v1/weather", h.getWeatherHandler(cfg.Weather))
	h.router.SERVE(cfg.Port)
}

// healthCheck endpoint for heartbeat or ping of the application server
func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{Status: http.StatusText(http.StatusOK)})
}
