package main

import "fmt"

func main() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	asterisk := '*'
	smile := '😀'
	acute := 'é'

	fmt.Printf("%v %v %v %v %v %v %v\n", pi, alpha, omega, bang, asterisk, smile, acute)
	fmt.Printf("%c%c%c%c %c %c %c\n", pi, alpha, omega, bang, asterisk, smile, acute)
}
