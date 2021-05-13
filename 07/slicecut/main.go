package main

import s "github.com/inancgumus/prettyslice"

func main() {
	// msg := []string{"h", "e", "l", "l", "o"}
	// fmt.Println(msg)
	// fmt.Println(msg[0:4])
	// fmt.Println(msg[:4])
	// fmt.Println(append(msg[:4], "!"))

	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"meroid", "defender", "rayman", "tempest",
		"ultima",
	}

	s.MaxPerLine = 4
	s.Show("items", items)
}
