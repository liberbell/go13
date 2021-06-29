package main

func main() {
	var (
		modydick  = book{title: "mody dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
		rubik     = puzzle{title: "rubik`s cube", price: 5}
	)

	var store list
	store = append(store, &minecraft, &tetris, modydick, rubik)
	store.print()

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
