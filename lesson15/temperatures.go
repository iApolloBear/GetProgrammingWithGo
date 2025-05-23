package main

import "fmt"

type (
	celsius    float64
	fahrenheit float64
	getRowFn   func(row int) (string, string)
)

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

const (
	rowFormat    = "| %8s | %8s |\n"
	numberFormat = "%.1f"
	line         = "======================="
)

func drawTable(hdr1 string, hdr2 string, getRow getRowFn) {
	fmt.Println(line)
	fmt.Printf(rowFormat, hdr1, hdr2)
	fmt.Println(line)
	for index := -40; index <= 100; index += 5 {
		row1, row2 := getRow(index)
		fmt.Printf(rowFormat, row1, row2)
	}
	fmt.Println(line)
}

func cToF(row int) (string, string) {
	c := fmt.Sprintf(numberFormat, celsius(row))
	f := fmt.Sprintf(numberFormat, celsius(row).fahrenheit())
	return c, f
}

func fToC(row int) (string, string) {
	f := fmt.Sprintf(numberFormat, fahrenheit(row))
	c := fmt.Sprintf(numberFormat, fahrenheit(row).celsius())
	return f, c
}

func main() {
	drawTable("ºC", "ºF", cToF)
	drawTable("ºF", "ºC", fToC)
}
