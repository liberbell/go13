package main

func main() {
	store := list{
		{"mody dick", 10, 118281600},
		{"odyssey", 10, "723622855"},
		{"hobbit", 25, nil},
		{"minecraft", 20},
		{"tetris", 5},
		{"rubik`s cube", 5},
		{"yoda", 150},
	}

	store.discount(.5)
	store.print()

}
