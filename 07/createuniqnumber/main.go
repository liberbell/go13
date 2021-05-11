package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const max = 5
	var uniqes [max]int

loop:
	for found := 0; found < max; found++ {
		n := rand.Intn(max)

		fmt.Print(n, " ")

		for _, u := range uniqes {
			if u == n {
				continue loop
			}
		}
	}
}
