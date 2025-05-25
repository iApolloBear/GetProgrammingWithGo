package main

import "fmt"

const line = "=========="

func getRow(row int, column int) rune {
	whitePieces := [8]rune{9820, 9822, 9821, 9819, 9818, 9821, 9822, 9820}
	blackPieces := [8]rune{9814, 9816, 9815, 9813, 9812, 9815, 9816, 9814}
	const whitePawn = 9823
	const blackPawn = 9817

	switch row {
	case 0:
		return blackPieces[column]
	case 1:
		return blackPawn
	case 6:
		return whitePawn
	case 7:
		return whitePieces[column]
	default:
		return 20
	}
}

func drawBoard() {
	var board [8][8]rune

	fmt.Println(line)
	for row := range board {
		fmt.Print("|")
		for column := range board[row] {
			fmt.Printf("%c", getRow(row, column))
		}
		fmt.Print("|\n")
	}
	fmt.Println(line)
}

func main() {
	drawBoard()
}
