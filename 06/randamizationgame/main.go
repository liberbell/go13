package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 10
	Usage    = `Welcome to the Luck Number game.
	The Program will pick %d random numbers.
	Your mission is to guess one these numbers.
	
	The greater your number is, harder it gets.
	
	Wanna plya?`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	if len(os.Args) != 2 {
		fmt.Println(Usage, maxTurns)
		return
	}
	guess, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("input a number.")
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
