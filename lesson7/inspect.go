package main

import "fmt"

func main() {
	year := 2018
	days := 365.2425
	text := "Ganges"
	isGood := true

	fmt.Printf("Type %T for %v\n", year, year)
	fmt.Printf("Type %T for %[1]v\n", days)
	fmt.Printf("Type %T for %[1]v\n", text)
	fmt.Printf("Type %T for %[1]v\n", isGood)
}
