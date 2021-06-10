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
	}
	fmt.Println(string(input))
}
