package main

import "fmt"

func main() {
	dwarfsArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	fmt.Printf("dwarfs: %T\n", dwarfs)
	fmt.Printf("dwarfsArray: %T\n", dwarfsArray)
}
