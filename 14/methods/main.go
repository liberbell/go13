package main

func main() {
	var (
		modydick  = book{title: "mody dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)

	game.print(minecraft)
	game.print(tetris)
	book.print(modydick)
}
