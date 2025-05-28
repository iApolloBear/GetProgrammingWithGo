package main

import "fmt"

func main() {
	slc := []int{}
	capacity := cap(slc)

	for i := range 100 {
		slc = append(slc, i)
		if capacity != cap(slc) {
			fmt.Printf("Last element before capacity is recalculated %v\n", i)
			capacity = cap(slc)
			fmt.Println(capacity)
		}
	}
}
