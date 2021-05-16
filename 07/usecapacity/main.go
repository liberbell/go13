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

	part = games[:0]
	s.Show("Part[:0]", part)
	s.Show("Part[:cap]", part[:4])
	s.Show("Part[:cap]", part[:cap(part)])

	for cap(part) != 0 {
		part = part[1:cap(part)]
		s.Show("part", part)
	}
}
