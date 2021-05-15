package main

import s "github.com/inancgumus/prettyslice"

func main() {
	s.PrintBacking = true

	var games []string
	s.Show("games", games)
}
