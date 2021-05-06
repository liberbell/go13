package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	positives := [3]string{"good!", "awesome!", "happy!"}
	negatives := [3]string{"bad.", "sad.", "terrible."}

	moods := [...][3]string{
		{"happy", "good", "awesome"},
		{"bad", "sad", "terrible"},
	}

	var mi int

	if len(os.Args) != 3 {
		fmt.Println("Usage: [your name] [positive|negative]")
		return
	}
	name, feeling := os.Args[1], os.Args[2]
	// feeling := os.Args[2]

	if feeling != "positive" {
		mi = 1
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods))

	if feeling != "positive" && feeling != "negative" {
		// fmt.Printf("%s %s\n", name, feeling)
		fmt.Println("Input positive/negative.")
	} else if feeling == "positive" {
		p := rand.Intn(len(positives))
		fmt.Printf("%s has a %s feelings.\n", name, positives[p])
	} else if feeling == "negative" {
		n := rand.Intn(len(negatives))
		fmt.Printf("%s has a %s feelings.\n", name, negatives[n])
	}

	fmt.Printf("%s feels %s\n", name, moods[mi][n])
}
