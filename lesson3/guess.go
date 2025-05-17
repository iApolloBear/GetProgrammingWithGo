package main

import (
	"fmt"
	"math/rand"
)

func main() {
	myNumber := 28
	for {
		magicNumber := rand.Intn(100) + 1
		fmt.Println(magicNumber)
		if magicNumber > myNumber {
			fmt.Println("too big.")
		}
		if magicNumber < myNumber {
			fmt.Println("too small.")
		}
		if magicNumber == myNumber {
			break
		}
	}
}
