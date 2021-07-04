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

}
