package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const maxTurns = 10

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	guess := args[0]

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess {
			fmt.Println("You win!")
			return
		}
	}
	fmt.Println("You lose.")
}
