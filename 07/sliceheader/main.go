package main

import s "github.com/inancgumus/prettyslice"

type collection [4]string

func main() {
	s.PrintElementAddr = true

	data := collection{"slices", "are", "awesome", "period"}
	change(data)
}

func change(data, collection) {
	var data collection
}
