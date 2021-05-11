package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	max, _ := strconv.Atoi(os.Args[1])

	// const max = 5
	var uniques []int

loop:
	for found := 0; found < max; {
		n := rand.Intn(max) + 1

		fmt.Print(n, " ")

		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}
		// uniques[found] = n
		uniques = append(uniques, n)
		found++
	}
	fmt.Println("\n\nuniques: ", uniques)
}