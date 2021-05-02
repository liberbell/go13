package main

import "fmt"

func main() {
	var books [1 + 3]string

	fmt.Printf("books: %T\n", books)
	fmt.Println("books: ", books)
	fmt.Printf("books: %q\n", books)
	fmt.Printf("books: %#v\n", books)
}
