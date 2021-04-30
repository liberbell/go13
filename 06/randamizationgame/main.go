package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxTurns = 3

func main() {
	rand.Seed(time.Now().UnixNano())
	guess := 10

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess {
			fmt.Println("You win!")
			return
		}
	}
	fmt.Println("You lose.")
}
