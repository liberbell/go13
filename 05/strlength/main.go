package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "漢字"
	fmt.Println(utf8.RuneCountInString(name))
}
