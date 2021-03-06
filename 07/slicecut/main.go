package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

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

	// s.MaxPerLine = 4
	// s.Show("items", items)

	// top3 := items[:3]
	// s.Show("top 3 items: ", top3)

	// last4 := items[9:]
	// s.Show("last 4 items: ", last4)

	// mid := last4[1:3]
	// s.Show("mid items: ", mid)

	const pageSize = 4
	l := len(items)

	for from := 0; from < l; from += pageSize {
		to := from + pageSize
		if to > l {
			to = l
		}

		// fmt.Printf("%d:%d\n", from, to)
		currentPage := items[from:to]

		head := fmt.Sprintf("Page #%d", (from/pageSize)+1)

		s.Show(head, currentPage)
	}
}
