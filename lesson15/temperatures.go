package main

import "fmt"

type (
	celsius       float64
	fahrenheit    float64
	temperatureFn func()
)

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func drawTable(printTemperature temperatureFn) {
	fmt.Println("=======================")
	printTemperature()
	fmt.Println("=======================")
}

func printCelsius() {
	fmt.Printf("| %-8v | %-8v |\n", "ºC", "ºF")
	fmt.Println("=======================")
	for index := -5; index <= 100; index += 5 {
		c := fmt.Sprintf("%vºC", celsius(index))
		f := fmt.Sprintf("%vºF", celsius(index).fahrenheit())
		fmt.Printf("| %-8v | %-8v |\n", c, f)
	}
}

func printFahrenheit() {
	fmt.Printf("| %-8v | %-8v |\n", "ºF", "ºC")
	fmt.Println("=======================")
	for index := -5; index <= 100; index += 5 {
		c := fmt.Sprintf("%vºF", fahrenheit(index))
		f := fmt.Sprintf("%.2fºC", fahrenheit(index).celsius())
		fmt.Printf("| %-8v | %-8v |\n", c, f)
	}
}

func main() {
	cel := printCelsius
	far := printFahrenheit
	drawTable(cel)
	drawTable(far)
}
