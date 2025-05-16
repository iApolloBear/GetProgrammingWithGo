package main

import "fmt"

func main() {
	const lightspeed = 28 // km/s
	distance := 96_300_000
	result := distance / lightspeed

	fmt.Println(result, "seconds")
	fmt.Println(result/86_400, "days")
	fmt.Println(96_300_000/100_800/24, "days")
}
