package main

import "fmt"

func main() {
	c := 'a'
	c += 3
	if c > 'z' {
		c -= 26
	}
	// fmt.Printf("%c\n", c)
	// x := 'g' - 'a' + 'A'
	// fmt.Printf("%c is %[1]v", x)

	// message := "L fdph, L vdz, L frqtxhuhg."
	//
	// for i := 0; i < 27; i++ {
	// 	letter := message[i] - 3
	// 	fmt.Printf("%c", letter)
	// }
}
