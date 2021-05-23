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

	var (
		px, py int
		cell   rune
	)

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	// board[0][0] = true
	buf := make([]rune, 0, width*height)

	board[px][py] = true

	// board[12][2] = true
	// board[16][2] = true
	// board[14][4] = true
	// board[10][6] = true
	// board[18][6] = true
	// board[12][7] = true
	// board[14][7] = true
	// board[16][7] = true

	// draw board
	for i := 0; i < 1000; i++ {
		buf = buf[:0]

		for y := range board[0] {
			for x := range board {
				cell = cellempty
				if board[x][y] {
					// fmt.Print("⚾️")
					cell = cellBall
				}
				// fmt.Print("X")
				// fmt.Print(string(cell))

				buf = append(buf, cell, ' ')
				// fmt.Print(" ")
			}
			buf = append(buf, '\n')
			// fmt.Println()
		}
		fmt.Print(string(buf))
	}
}
