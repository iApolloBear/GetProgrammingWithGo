package main

import "fmt"

type (
	celsius    float64
	kelvin     float64
	fahrenheit float64
)

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func main() {
	var k kelvin = 294.0

	c := k.celsius()
	fmt.Println(c)
}
