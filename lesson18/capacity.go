package main

import "fmt"

func main() {
	slc := []int{}
	capacity := cap(slc)

	for i := range 100 {
		slc = append(slc, i)
		if capacity != cap(slc) {
			capacity = cap(slc)
			fmt.Println(capacity)
		}
	}
}
