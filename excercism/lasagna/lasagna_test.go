package lasagna

import (
	"testing"
)

func TestOvenTime(t *testing.T) {
	expected := 40
	if OvenTime != expected {
		t.Fatalf(`OventTime = %q, expected %#q, nil`, OvenTime, expected)
	}
}

func TestRemainingOvenTime(t *testing.T) {
	actual := 30
	expected := 10
	result := RemainingOvenTime(actual)
	if result != expected {
		t.Fatalf(`RemainingOvenTime(%q) = %q, expected %#q, nil`, actual, result, expected)
	}
}

func TestPreparationTimeTest(t *testing.T) {
	numberOfLayers := 2
	expected := 4
	result := PreparationTime(numberOfLayers)
	if result != expected {
		t.Fatalf(`PreparationTime(%q) = %q, expected %#q, nil`, numberOfLayers, result, expected)
	}
}

func TestElapsedTimeTest(t *testing.T) {
	var remainingTime int = RemainingOvenTime(30)
	var numberOfLayers int = PreparationTime(2)
	var elapsedTime int = ElapsedTime(numberOfLayers, remainingTime)
	var expected int = 14
	if elapsedTime != expected {
		t.Fatalf(`ElapsedTime(%q, %q) = %q, expected %#q, nil`, numberOfLayers, remainingTime, elapsedTime, expected)
	}
}
