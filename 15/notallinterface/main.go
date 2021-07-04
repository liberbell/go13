package main

func main() {
	l := list{
		{title: "mody dick", price: 10, released: 118281600},
		{title: "odyssey", price: 10, released: "723622855"},
		{title: "hobbit", price: 25},
		{title: "minecraft", price: 20},
		{title: "tetris", price: 5},
		{title: "rubik`s cube", price: 5},
		{title: "yoda", price: 150},
	}

	l.discount(.5)
	l.print()

}
