package main

import s "github.com/inancgumus/prettyslice"

func main() {
	s.PrintBacking = true

	var games []string
	s.Show("games", games)

	games = []string{}
	s.Show("games", games)

	games = []string{"pacman", "mario", "tetoris", "doom"}
	s.Show("games", games)
	s.Show("games", []string{"one", "two"})

	part := games
	s.Show("Part", part)
}
