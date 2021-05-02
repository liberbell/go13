package main

import "fmt"

func main() {
	var books [4]string

	fmt.Printf("books: %T\n", books)
	fmt.Println("books: ", books)
	fmt.Printf("books: %q\n", books)
}
