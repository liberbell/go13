package main

import "fmt"

// const ok = true

func nope() {
	const ok = true
	var hello = "HELLO"
}

func main() {
	// fmt.Println(("Hello World"))
	// var hello = "Hello"
	fmt.Println(hello, ok)
}
