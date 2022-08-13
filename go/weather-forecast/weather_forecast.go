// Package weather is being used to advertise weather conditions for a given location.
package weather

// CurrentCondition contains the current weather condition.
var CurrentCondition string

// CurrentLocation represents the interested current location.
var CurrentLocation string

// Forecast is a method which combines a current location, an existing condition and advertise the weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
