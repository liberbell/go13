package main

import "fmt"

func main() {
	var book [5]string
	book[0] = "doracula"
	book[1] = "1984"
	book[2] = "island"

	games := []string{"pokemon", "sims"}
	// game[0] = "pokemon"
	// game[1] = "sims"
	newGame := []string{"pacman", "doom", "pong"}
	newGame = games

	var ok string
	for i, game := range games {
		if game := newGame[i] {
			ok = "not"
			break
		}
	}
	fmt.Printf("games and newGame are %sequal\n\n", ok)

	newBooks := [5]string{"ulysses", "fire"}
	book = newBooks

	fmt.Printf("books     : %#v\n", book)
	fmt.Printf("newBooks  : %#v\n", newBooks)
	fmt.Printf("game      : %#v\n", games)
	fmt.Printf("game      : %T\n", games)

	fmt.Printf("game len  : %d\n", len(games))
	fmt.Printf("nil?      : %t\n", games == nil)
}
