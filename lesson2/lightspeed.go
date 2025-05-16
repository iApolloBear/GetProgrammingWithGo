package main

import "fmt"

func main() {
	distance, speed := 56_000_000, 100_800
	const hoursInADay, minutesPerHour = 24, 60

	fmt.Println(distance/speed, "seconds")
}
