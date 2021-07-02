package main

func main() {
	store := list{
		book{title: "mody dick", price: 10, published: 118281600},
		book{title: "odyssey", price: 10, published: "723340855"},
		book{title: "hobbit", price: 10, published: 25},
		&game{title: "minecraft", price: 20},
		&game{title: "tetris", price: 5},
		puzzle{title: "rubik`s cube", price: 5},
		&toy{title: "yoda", price: 150},
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
