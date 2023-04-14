package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CalculateWorkingCarsPerHour(numberOfCars int, percent float64) float64 {
	var average float64 = 0.00
	average = float64(numberOfCars) * (percent / 100)
	return average
}

func CalculateWorkingCarsPerMinute(numberOfCars int, percent float64) int {
	var average int = 0
	var average_per_hour float64 = CalculateWorkingCarsPerHour(numberOfCars, percent)
	average = int(average_per_hour / 60)
	return average
}

const Cost = 10000
const SalePrice = 95000

func CalculateCost(numberOfCars int) uint {
	var final_cost uint = 0
	var number string = strconv.Itoa(numberOfCars)
	var new_number string = ""
	for _, char := range number {
		n, _ := strconv.Atoi(string(char))
		new_number += fmt.Sprintf("%d,", n)
	}
	var new_string []string = strings.SplitN(new_number, ",", 2)
	nn1, _ := strconv.Atoi(new_string[0])
	nn2, _ := strconv.Atoi(strings.Trim(new_string[1], ","))
	final_cost = uint(nn1)*SalePrice + uint(nn2)*Cost
	return final_cost
}

func main() {
	fmt.Printf("%f\n", CalculateWorkingCarsPerHour(1547, 90))
	fmt.Printf("%d\n", CalculateWorkingCarsPerMinute(1105, 90))
	fmt.Printf("%d\n", CalculateCost(100))
	fmt.Printf("%d\n", CalculateCost(21))
}
