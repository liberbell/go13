package main

import (
	"fmt"
)

func main() {
	text := `So i come from Python, and they encourages you to program with readable not-that-
I’m 出力は5V/3A、9V/3A、15V/3A、20V/2.25A、PPS出力が3.3～16V=3A、3.3～21V=2.25Aなどとなっている。`

	for _, r := range text {
		// r, size := utf8.DecodeRuneInString(text[i:])
		fmt.Printf("%c", r)

		// i += size
	}
	fmt.Println()
}
