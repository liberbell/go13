package main

import (
	"bufio"
	"os"
)

var user struct {
	Name        string
	Permissions map[string]bool
}

func main() {
	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
		// input = append(input, '\n')
	}
	// fmt.Println(string(input))
	var users []user
}
