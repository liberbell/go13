package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("Usage: [yuar name]")
		return
	}
	name := arg[0]
	moods := [...]string{
		"happy", "good", "awesome",
		"sad", "bad", "terrible",
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods))
}
