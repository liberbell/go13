package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	positives := [3]string{"good!", "awesome!", "happy!"}
	// negatives := [3]string{"bad.", "happy.", "terrible."}

	if len(os.Args) != 3 {
		fmt.Println("Usage: [name][positive|negative]")
		return
	}
	name := os.Args[1]
	feeling := os.Args[2]

	rand.Seed(time.Now().UnixNano())

	if feeling == "positive" {
		fmt.Printf("%s %s\n", name, feeling)
		i := rand.Intn(len(positives))
		fmt.Printf("%s have a %s feelings.\n", name, positives[i])
	}
}
