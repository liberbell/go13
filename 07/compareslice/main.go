package main

import "fmt"

func main() {
	var book [5]string
	book[0] = "doracula"
	book[1] = "1984"
	book[2] = "island"

	fmt.Printf("books     : %#v\n", book)
}
