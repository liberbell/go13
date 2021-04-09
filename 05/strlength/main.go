package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	name := "漢字"
	fmt.Println(utf8.RuneCountInString(name))

	msg := os.Args[1]
	l := len(msg)
	s := msg + strings.Repeat("!", l)
	s = strings.ToUpper(s)
	fmt.Println(s)
}
