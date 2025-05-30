package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km/s
	const secondsPerDay = 86400

	var distance int64 = 41.3e12 // 41_300_000_000_000
	fmt.Println("Alpha Centauri is", distance, "km away.")

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed.")
	distance = 56e6 // 56 + six zeroes = 56_000_000
	fmt.Println(distance)
	distance = 401e6 // 401 + six zeroes = 401_000_000
	fmt.Println(distance)
}
