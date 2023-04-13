package lasagna

import (
	"math"
)

const OvenTime = 40

func RemainingOvenTime(actual int) int {
	var remaingTime float64 = math.Abs(float64(actual - OvenTime))
	return int(remaingTime)
}

func PreparationTime(numberOfLayers int) int {
	var layers int = 2
	return layers * numberOfLayers
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	var elapseTime int = numberOfLayers + actualMinutesInOven
	return elapseTime
}
