package main

import f "fmt"

const ok = true

// func nope() {
// 	const ok = true
// 	var hello = "HELLO"
// 	_ = hello
// }

func main() {
	f.Println(("Hello World"))
	// var hello = "Hello"
	// fmt.Println(hello, ok)
	// bye()
	// hey()
}

func bye() {
	f.Println("Bye")
}
