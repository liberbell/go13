package main

func newParser() parser {
	return parser{sum: make(map[string]result)}
}
