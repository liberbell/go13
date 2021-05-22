package main

import (
	"fmt"
)

func main() {
	const (
		width  = 50
		height = 15
	)

	var cell rune

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}
	board[0][0] = true

	// draw board
	for y := range board[0] {
		for x := range board {
			cell = ' '
			if board[x][y] {
				// fmt.Print("⚾️")
				cell = '⚾️'
			}
			// fmt.Print("X")
			fmt.Print(cell)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
