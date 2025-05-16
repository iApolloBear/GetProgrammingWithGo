package main

import "fmt"

func main() {
	const days = 28
	const distance = 56_000_000
	const kmsPerDay = distance / days
	const kmsPerHour = kmsPerDay / 24
	fmt.Printf("%v km/h", kmsPerHour)
}
