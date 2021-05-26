package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	text := `
出力は5V/3A、9V/3A、15V/3A、20V/2.25A、PPS出力が3.3～16V=3A、3.3～21V=2.25Aなどとなっている。`

	for i := 0; i < len(text); i++ {
		r, _ := utf8.DecodeRuneInString(text)
		fmt.Printf("%c", r)
	}

}
