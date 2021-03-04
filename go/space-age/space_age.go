package space

import "math"

// Planet is a custom type for our planets
type Planet string

// planets is a map with Planet name and its period
var planets = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

var secondsInAYearOnEarth = 31557600

// Age is a method returning the age of a given Planet
func Age(seconds float64, planet Planet) float64 {
	result := seconds / (planets[planet] * float64(secondsInAYearOnEarth))
	return math.Round(result*100) / 100
}
