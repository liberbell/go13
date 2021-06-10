package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
		input = append(input, '\n')
	}
	fmt.Println(string(input))
}
