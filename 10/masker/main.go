package main

import (
	"fmt"
	"os"
)

// Get and check the input
// Create a byte buffer and use it as the output
// Write input to the buffer as it is and print it
// Mask the link

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("gimme something to mask!")
		return
	}

	const (
		link  = "http://"
		nlink = len(link)
		mask  = '*'
	)

	var (
		text = args[0]
		size = len(text)
		buf  = make([]byte, 0, size)

		in bool
	)

	// fmt.Println("text size: ", size)

	for i := 0; i < size; i++ {
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			// fmt.Printf(`text[%d : %[1]d+%d] = `, i, nlink)
			// fmt.Println(text[i : i+nlink])
			in = true
		}

		c := text[i]

		switch c {
		case ' ', '\t', '\n':
			in = false
		}

		if in {
			c = mask
		}

		buf = append(buf, c)
	}
	fmt.Println(string(buf))
}
