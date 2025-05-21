package main

import "fmt"

// kelvinToCelsius converts K to C
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32
}

func kelvinToFahrenheit(k float64) float64 {
	// return celsiusToFahrenheit(kelvinToCelsius(k))
	k -= 459.67
	return k
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "K is ", celsius, "C\n")
	fahrenheit := celsiusToFahrenheit(29.0)
	fmt.Print(29.0, "C is ", fahrenheit, "F\n")
	fmt.Print(0, "K is ", kelvinToFahrenheit(0), "F\n")
}
