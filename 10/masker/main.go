package main

import (
	"fmt"
	"os"
)

// Get and check the input
// Create a byte buffer and use it as the output

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("gimme something to mask!")
		return
	}

	var (
		text = args[0]
		size = len(text)
		buf  = make([]byte, 0, size)
	)

	for i := 0; i < size; i++ {
		buf = append(buf, text[1])
		fmt.Println(string(buf))
	}
}
