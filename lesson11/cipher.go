package main

import (
	"fmt"
	"strings"
)

func main() {
	plainText := "your message goes here"
	keyword := "GOLANG"
	formattedText := strings.ToUpper(strings.Replace(plainText, " ", "", 3))

	for i := 0; i < len(formattedText); i++ {
		letterToEncrypt := formattedText[i] - 'A'
		keyWordLetter := keyword[i%len(keyword)] - 'A'
		ciphered := (letterToEncrypt + keyWordLetter) % 26
		fmt.Printf("%c", ciphered+'A')
	}
}
