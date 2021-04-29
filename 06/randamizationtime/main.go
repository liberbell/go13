package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	unixt := time.Now()
	// fmt.Println(unixt)

	rand.Seed(unixt.UnixNano())

	guess := 10
	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}

}
