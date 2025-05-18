package main

import "fmt"

func main() {
	piggyBank := 0.0
	for count := 0; count < 11; count++ {
		piggyBank += 0.1
	}
	fmt.Println(piggyBank)
}
