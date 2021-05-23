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

	// board[0][0] = true
	buf := make([]rune, 0, width*height)

	board[12][2] = true
	board[16][2] = true
	board[14][4] = true
	board[10][6] = true
	board[18][6] = true
	board[12][7] = true
	board[14][7] = true
	board[16][7] = true

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
