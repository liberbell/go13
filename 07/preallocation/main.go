package main

import "fmt"

func main() {
	s := make([]int, 0, 5)
	fmt.Println(s)
	s = append(s, 42)
	fmt.Println(s)
}
