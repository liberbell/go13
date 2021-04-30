package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const maxTurns = 10

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	guess, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("input one number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pickup a positive number.")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess {
			fmt.Println("You win!")
			return
		}
	}
	fmt.Println("You lose.")
}
