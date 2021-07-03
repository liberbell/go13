package main

func main() {
	store := list{
		book{"mody dick", 10, published: 118281600},
		book{"odyssey", 10, published: "723622855"},
		book{"hobbit", 25},
		&game{"minecraft", 20},
		&game{"tetris", 5},
		puzzle{"rubik`s cube", 5},
		&toy{"yoda", 150},
	}

	store.discount(.5)

	store.print()

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
