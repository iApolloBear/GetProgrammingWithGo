package main

import (
	"fmt"
	"math/rand"
)

func main() {
	piggyBank := 0.0
	for piggyBank <= 20.0 {
		switch coin := rand.Intn(3); coin {
		case 0:
			piggyBank += 0.05
		case 1:
			piggyBank += 0.10
		default:
			piggyBank += 0.25
		}
	}
	fmt.Printf("%5.2f", piggyBank)
}
