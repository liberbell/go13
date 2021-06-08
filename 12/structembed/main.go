package main

import "fmt"

func main() {
	type text struct {
		title string
		words int
	}

	type book struct {
		title string
		words int
		isbn  string
	}

	moby := book{title: "mody dick", words: 206052, isbn: "102030"}
	fmt.Printf("%s has %d words (isbn: %s)\n", moby.title, moby.words, moby.isbn)
}
