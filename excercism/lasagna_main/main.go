package main

import (
	"excercism/lasagna"
	"fmt"
)

func main() {
	var remainingTime int = lasagna.RemainingOvenTime(30)
	fmt.Printf("RemainingOvenTime: %d\n", remainingTime)
	var numberOfLayers int = lasagna.PreparationTime(2)
	fmt.Printf("numberOfLayers: %d\n", numberOfLayers)
	var elapsedTime int = lasagna.ElapsedTime(numberOfLayers, remainingTime)
	fmt.Printf("elapseTime: %d\n", elapsedTime)

}
