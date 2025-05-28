package main

import (
	"fmt"
	"strings"
)

type Planets []string

func (planets Planets) terraform() {
	for index := range planets {
		planets[index] = "New " + planets[index]
	}
}

func main() {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	Planets(planets[3:4]).terraform()
	Planets(planets[6:]).terraform()
	fmt.Println(strings.Join(planets, ", "))
}
