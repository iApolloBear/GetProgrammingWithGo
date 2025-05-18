package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("%-17v%-5v%-11v%-5v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("======================================")

	for range 10 {
		spaceline := ""
		tripType := ""
		speed := rand.Intn(15) + 16 // Speed in km/s
		kmPerDay := speed * 60 * 60 * 24
		distance := 62_100_000
		days := distance / kmPerDay
		price := speed + 20

		if rand.Intn(2) == 0 {
			tripType = "One Way"
		} else {
			tripType = "Round-trip"
			price *= 2
		}

		switch spacelineIndx := rand.Intn(3); spacelineIndx {
		case 0:
			spaceline = "Virgin Galactic"
		case 1:
			spaceline = "SpaceX"
		default:
			spaceline = "Space Adventures"
		}

		fmt.Printf("%-17v%4v %-10v $%4v\n", spaceline, days, tripType, price)
	}
}
