package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	guess := 10

	for turn := 0; turn < 5; turn++ {
		n := rand.Intn(guess + 1)
	}
}
