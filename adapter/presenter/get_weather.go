// Package presenter returns the response from the use case logic to the API caller
package presenter

import (
	"encoding/json"

	"github.com/blevels/weatherAPI/domain/entity"
	"github.com/blevels/weatherAPI/usecase"
)

type GetWeatherPresenter struct{}

// NewGetWeatherPresenter creates new GetWeatherPresenter interface used to present data to the API caller
func NewGetWeatherPresenter() usecase.GetWeatherPresenter {
	return GetWeatherPresenter{}
}

// Output actually returns the use case response to the API caller
func (g GetWeatherPresenter) Output(w map[string]interface{}) usecase.GetWeatherOutput {
	var result entity.WeatherOutput

	DecodeViaJSON(w, &result)

	return usecase.GetWeatherOutput{
		Weather: result,
	}
}

// DecodeViaJSON takes the map data and passes it through encoding/json to convert it into the
// given Go native structure pointed to by v. v must be a pointer to a struct.
func DecodeViaJSON(data interface{}, v interface{}) error {
	// Perform the task by simply marshalling the input into JSON, then unmarshalling
	// it into target native Go struct.
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}
