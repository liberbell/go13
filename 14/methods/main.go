package main

func main() {
	store := list{
		&book{product{"mody dick", 10}, 118281600},
		&book{product{"odyssey", 10}, "723622855"},
		&book{product{"hobbit", 25}, nil},
		&game{product{"minecraft", 20}},
		&game{product{"tetris", 5}},
		&puzzle{product{"rubik`s cube", 5}},
		&toy{product{"yoda", 150}},
	}

	store.discount(.5)

	store.print()

	// t := &toy{product{"yoda", 150}}
	// fmt.Printf("%#v\n", t)

	// b := &book{product{"mody dick", 10}, 118281600}
	// fmt.Printf("%#v\n", b)

	// var p printer

	// p = modydick
	// p = rubik
	// p = &minecraft
	// p = &tetris
	// tetris.discount(.5)
	// p.print()

	// fmt.Println(store[0] == &minecraft)
	// fmt.Println(store[3] == rubik)

	// var items []*game
	// items = append(items, &minecraft, &tetris)

	// my := list(items)
	// // my = nil
	// my.print()

	// minecraft.discount(.5)

	// modydick.print()
	// minecraft.print()
	// tetris.print()

	// var h huge

	// for i := 0; i < 10; i++ {
	// 	h.addr()
	// }
}
