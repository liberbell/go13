package main

import (
	"fmt"
)

func main() {
	const (
		width  = 50
		height = 15

		cellempty = ' '
		cellBall  = '⚾'
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
			cell = cellempty
			if board[x][y] {
				// fmt.Print("⚾️")
				cell = cellBall
			}
			// fmt.Print("X")
			fmt.Print(string(cell))
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
