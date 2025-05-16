package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const min = 56_000_000
	const max = 401_000_000
	const distance = max - min
	fmt.Println(rand.Intn(distance) + min)
}
