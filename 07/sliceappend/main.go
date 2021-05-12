package main

import s "github.com/inancgumus/prettyslice"

func main() {
	var todo []string

	todo = append(todo, "sing")

	s.Show("todo:", todo)
}
