package main

import (
	"fmt"
)

func main() {
	const (
		width  = 50
		height = 15
	)

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}
	board[0][0] = true

	// draw board
	for y := range board[0] {
		for x := range board {
			fmt.Print("X")
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
