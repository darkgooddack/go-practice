package main

import "fmt"

func createTemperatureAdjuster() (func(change float64) (float64, float64, float64), float64) {
	baseTemperature := 90.0

	adjustTemperature := func(change float64) (float64, float64, float64) {
		oldTemperature := baseTemperature
		baseTemperature = baseTemperature + change
		return oldTemperature, change, baseTemperature
	}

	return adjustTemperature, baseTemperature
}

func main() {
	adjustTemp, originTemp := createTemperatureAdjuster()
	fmt.Printf("Original temperature: %.2f\n", originTemp)

	oldTemp, change, newTemp := adjustTemp(1.5)
	fmt.Printf("Base temperature: %.2f + %.2f = %.2f\n", oldTemp, change, newTemp)

	oldTemp, change, newTemp = adjustTemp(-2.5)
	fmt.Printf("Base temperature: %.2f + %.2f = %.2f\n", oldTemp, change, newTemp)
}
