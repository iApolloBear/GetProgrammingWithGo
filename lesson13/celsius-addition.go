package main

import "fmt"

func main() {
	type celsius float64
	type fahrenheit float64
	const degrees = 20
	var warmUp float64 = 10
	var temperature celsius = degrees
	temperature += celsius(warmUp)
	fmt.Println(temperature)
}
