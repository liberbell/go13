package main

import "fmt"

func main() {
	msg := []string{"h", "e", "l", "l", "o"}
	fmt.Println(msg)
	fmt.Println(msg[0:4])
	fmt.Println(msg[:4])
	fmt.Println(append(msg[:4], "!"))
}
