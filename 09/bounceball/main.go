package main

func main() {
	const (
		width  = 50
		height = 10
	)

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}
}
