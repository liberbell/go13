package main

import "fmt"

func main() {
	var book [5]string
	book[0] = "doracula"
	book[1] = "1984"
	book[2] = "island"

	var game []string

	fmt.Printf("books     : %#v\n", book)
	fmt.Printf("game      : %#v\n", game)
	fmt.Printf("game      : %T\n", game)
	fmt.Printf("game len  : %d\n", len(game))
}
