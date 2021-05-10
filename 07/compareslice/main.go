package main

import "fmt"

func main() {
	var book [5]string
	book[0] = "doracula"
	book[1] = "1984"
	book[2] = "island"

	// var game []string
	game := []string{"pokemon", "sims"}
	// game[0] = "pokemon"
	// game[1] = "sims"

	newBooks := [5]string{"ulysses", "fire"}
	book = newBooks

	fmt.Printf("books     : %#v\n", book)
	fmt.Printf("newBooks  : %#v\n", newBooks)
	fmt.Printf("game      : %#v\n", game)
	fmt.Printf("game      : %T\n", game)

	fmt.Printf("game len  : %d\n", len(game))
	fmt.Printf("nil?      : %t\n", game == nil)
}
