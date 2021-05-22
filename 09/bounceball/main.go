package main

import (
	"fmt"
)

func main() {
	const (
		width  = 50
		height = 10
	)

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}
	// draw board
	for range board[0] {
		for range board {
			fmt.Print("X")
		}
		fmt.Println()
	}
}
