// Package usecase implements the applications business/domain logic. Each logical grouping requires its own file and instrumentation
package usecase

import (
	"context"
	"time"

	"github.com/blevels/weatherAPI/domain/entity"
)

type (
	// ApiRequestSender port
	ApiRequestSender interface {
		Send(context.Context, entity.Weather) (map[string]interface{}, error)
	}

	// GetWeatherUseCase Input port
	GetWeatherUseCase interface {
		Execute(context.Context, GetWeatherInput) (GetWeatherOutput, error)
	}

	// GetWeatherPresenter Presenter/Output port sends data to the caller
	GetWeatherPresenter interface {
		Output(map[string]interface{}) GetWeatherOutput
	}

	// GetWeatherInput Input data received by the API
	GetWeatherInput struct {
		Longitude string `json:"longitude"`
		Latitude  string `json:"latitude"`
	}

	// GetWeatherOutput Output data format
	GetWeatherOutput struct {
		Weather entity.WeatherOutput
	}

	// GetWeatherInteractor Provides the interfaces between the external layers of the application and the inner layers
	GetWeatherInteractor struct {
		pre       GetWeatherPresenter
		requester ApiRequestSender
	}
)

// NewGetWeatherInteractor creates new getWeatherInteractor with its dependencies injected
func NewGetWeatherInteractor(
	requester ApiRequestSender,
	pre GetWeatherPresenter,
) GetWeatherUseCase {
	return GetWeatherInteractor{
		requester: requester,
		pre:       pre,
	}
}

// Execute orchestrates the use case for the domain logic
func (g GetWeatherInteractor) Execute(ctx context.Context, i GetWeatherInput) (GetWeatherOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := g.requester.Send(ctx, entity.Weather(i))
	if err != nil {
		return g.pre.Output(map[string]interface{}{}), err
	}

	return g.pre.Output(res), nil
}
