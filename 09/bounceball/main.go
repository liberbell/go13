package main

import (
	"fmt"
	"screen"
	"time"
)

func main() {
	const (
		width  = 50
		height = 15

		cellEmpty = ' '
		cellBall  = '⚾'
		maxFrames = 1200

		speed = time.Second / 20
	)

	var (
		px, py int
		vx, vy = 1, 1
		cell   rune
	)

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	// board[0][0] = true
	buf := make([]rune, 0, width*height)
	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

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
		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
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
		screen.MoveTopLeft()
		fmt.Print(string(buf))

		time.Sleep(speed)
	}
}
