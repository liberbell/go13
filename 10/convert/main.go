package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "あいうえお 🙉"

	bytes := []byte(str)
	// bytes[0] = 'N'
	// bytes[1] = 'O'

	// str = string(bytes)
	str = string(bytes)

	fmt.Printf("%s\n", str)
	fmt.Printf("\t%d bytes\n", len(str))
	fmt.Printf("\t%d Runes \n", utf8.RuneCountInString(str))

	fmt.Printf("% x\n", bytes)
	fmt.Printf("\t%d bytes\n", len(bytes))
	fmt.Printf("\t%d Runes \n", utf8.RuneCount(bytes))

	fmt.Println()
	for i := range str {
		fmt.Printf("str[%2d] = %-2x\n", i, str[i])
	}
}
