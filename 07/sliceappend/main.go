package main

import s "github.com/inancgumus/prettyslice"

func main() {
	var todo []string

	todo = append(todo, "sing")
	todo = append(todo, "run", "code", "play")
	todo = append(todo, "code")
	todo = append(todo, "play")

	s.Show("todo:", todo)
}
