package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var piggyBank uint16 = 0

	for piggyBank <= 2000 {
		switch coin := rand.Intn(3); coin {
		case 0:
			piggyBank += 5
		case 1:
			piggyBank += 10
		case 2:
			piggyBank += 25
		}
	}

	var remainder uint16 = piggyBank % 100
	fmt.Printf("$%v.%02v\n", piggyBank/100, remainder)
}
