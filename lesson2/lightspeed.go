package main

import "fmt"

func main() {
	const lightspeed = 299792
	distance := 56_000_000

	fmt.Println(distance/lightspeed, "seconds")
	distance = 401_000_000
	fmt.Println(distance/lightspeed, "seconds")
}
