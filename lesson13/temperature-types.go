package main

import "fmt"

type (
	celsius float64
	kelvin  float64
)

// kelvinToCelsius converts K to C
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Println(k, "K is", c, "C")
	var cels celsius = 127.0
	kelv := celsiusToKelvin(cels)
	fmt.Print(cels, " C is ", kelv, " K")
}
