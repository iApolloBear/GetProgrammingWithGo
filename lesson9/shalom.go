package main

import "fmt"

func main() {
	shalom := "shalom"

	for i := 0; i < 6; i++ {
		fmt.Printf("%c = %[1]v\n", shalom[i])
	}
}
