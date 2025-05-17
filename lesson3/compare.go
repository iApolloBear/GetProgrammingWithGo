package main

import (
	"fmt"
)

func main() {
	fmt.Println("There is a sign near the entrance that reads 'No Minors'.")
	age := 41
	minor := age < 18
	fmt.Printf("At age %v, im I a minor? %v\n", age, minor)
	fmt.Println("apple" > "banana")
}
