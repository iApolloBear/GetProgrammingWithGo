package main

import (
	"fmt"
	"strings"
)

type Planets []string

func (p Planets) terraform() {
	for index, planet := range p {
		p[index] = "New " + planet
	}
}

func main() {
	planets := []string{"Mars", "Uranus", "Neptune"}
	Planets(planets).terraform()
	fmt.Println(strings.Join(planets, ", "))
}
