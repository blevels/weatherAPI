// Package presenter returns the response from the use case logic to the API caller
package presenter

import (
	"github.com/blevels/weatherAPI/domain/entity"
	"github.com/blevels/weatherAPI/usecase"
)

type GetWeatherPresenter struct{}

// NewGetWeatherPresenter creates new GetWeatherPresenter interface used to present data to the API caller
func NewGetWeatherPresenter() usecase.GetWeatherPresenter {
	return GetWeatherPresenter{}
}

// Output actually returns the use case response to the API caller
func (g GetWeatherPresenter) Output(w entity.Weather) usecase.GetWeatherOutput {
	return usecase.GetWeatherOutput{
		Weather: map[string]interface{}{},
	}
}
