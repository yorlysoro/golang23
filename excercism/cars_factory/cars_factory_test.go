package main

import (
	"testing"
)

func TestCalculateWorkingCarsPerHour(t *testing.T) {
	var expected float64 = 1392.3
	numberOfCars := 1547
	percent := 90
	result := CalculateWorkingCarsPerHour(numberOfCars, percent)
	if result != expected {
		t.Fatalf(`CalculateWorkingCarsPerHour(%q, %q) = %f, expected %f, nil`, numberOfCars, percent, result, expected)
	}
}

func TestCalculateWorkingCarsPerMinute(t *testing.T) {
	var expected int = 16
	numberOfCars := 1105
	percent := 90
	result := CalculateWorkingCarsPerMinute(numberOfCars, percent)
	if result != expected {
		t.Fatalf(`CalculateWorkingCarsPerMinute(%d, %d) = %d, expected %d, nil`, numberOfCars, percent, result, expected)
	}

}

func TestCalculateCost(t *testing.T) {
	var expected uint = 355000
	var numberOfCars uint = 37
	var result uint = CalculateCost(numberOfCars)
	if result != expected {
		t.Fatalf(`CalculateWorkingCarsPerMinute(%d) = %d, expected %d, nil`, numberOfCars, result, expected)
	}
	expected = 200000
	numberOfCars = 21
	result = CalculateCost(numberOfCars)
	if result != expected {
		t.Fatalf(`CalculateWorkingCarsPerMinute(%d) = %d, expected %d, nil`, numberOfCars, result, expected)
	}
}
