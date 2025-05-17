package main

import "fmt"

func main() {
	room := "cave"

	if room == "cave" {
		fmt.Println("You found yourself in a dimly lit cavern.")
	} else if room == "entrance" {
		fmt.Println("There is a cavern entrance here and a path to the east.")
	} else if room == "mountain" {
		fmt.Println("There is a cliff here. A path leads west down the mountains.")
	} else {
		fmt.Println("Everything is white.")
	}
}
