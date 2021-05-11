package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const max = 5

	for found := 0; found < max; found++ {
		n := rand.Intn(max)
	}
}
