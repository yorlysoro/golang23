// Package weather_forecast.go calculates the forecast weather in Goblinocus a small city in Neverland.
package main

import (
	"fmt"
)

// CurrentCondition The current weather condition in Goblinocus.
var CurrentCondition string

// CurrentLocation The current location in Goblinocus.
var CurrentLocation string

// Forecast function return the current forecast in  a determinate location, received tow string in the parameters (city and condition) both string, and return a string indicate a current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

func main() {
	fmt.Printf(Forecast("Goblanicous", "34C"))
}
