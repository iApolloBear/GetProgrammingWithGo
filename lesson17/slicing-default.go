package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:]
	allPlanets := planets[:]
	colonized := terrestrial[2:]
	fmt.Println(terrestrial, gasGiants, iceGiants, allPlanets)
	fmt.Println(colonized)
}
