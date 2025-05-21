package main

import "fmt"

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"

	for i := 0; i < len(cipherText); i++ {
		cipherLetter := cipherText[i] - 'A'
		keywordLetter := keyword[i%len(keyword)] - 'A'
		deciphered := (cipherLetter - keywordLetter + 26) % 26
		fmt.Printf("%c", deciphered+'A')
	}
}
