package main

func main() {
	var (
		modydick  = book{title: "mody dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
		rubik     = puzzle{title: "rubik`s cube", price: 5}
		// yoda      = toy{title: "yoda", price: 150}
	)

	var store list
	store = append(store, &minecraft, &tetris, modydick, rubik, &yoda)
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
