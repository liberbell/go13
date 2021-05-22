package main

func main() {
	board := make([][]bool, 50)
	for row := range board {
		board[row] = make([]bool, 10)
	}
}
