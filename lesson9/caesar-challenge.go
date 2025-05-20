package main

import "fmt"

func main() {
	phrase := "L fdph, L vdz, L frqtxhuhg."
	for i := 0; i < len(phrase); i++ {
		c := phrase[i]
		if c >= 'a' && c <= 'b' {
			c += 23
		}
		c -= 3
		fmt.Printf("%c", c)
	}
}
