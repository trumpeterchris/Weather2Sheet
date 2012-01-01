package forecast

import (
	"fmt"
)

type ForecastConditions struct {
	Forecast Forecast
}

type Forecast struct {
	Txt_forecast Txt_forecast
}

type Txt_forecast struct {
	Date        string
	Forecastday []Forecastday
}

type Forecastday struct {
	Title   string
	Fcttext string
}

// printForecast prints the forecast for a given station to standard out
func PrintForecast(obs *ForecastConditions, stationId string) {
	t := obs.Forecast.Txt_forecast
	fmt.Printf("Issued at %s\n", t.Date)
	for _, f := range t.Forecastday {
		fmt.Printf("%s: %s\n", f.Title, f.Fcttext)
	}
}
