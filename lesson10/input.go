package main

import "fmt"

func main() {
	result := false
	switch word := "true"; word {
	case "true", "yes", "1":
		result = true
	case "false", "no", "0":
		result = false
	default:
		fmt.Println("Error: not valid word")
	}
	fmt.Println(result)
}
