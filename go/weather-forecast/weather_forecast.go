// Package weather provides a weather forecast.
package weather

// CurrentCondition is the most recent condition forecast.
var CurrentCondition string

// CurrentLocation is the most recent city to receive a forecast.
var CurrentLocation string

// Forecast takes a city and a condition and returns a forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
