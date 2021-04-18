package main

import "fmt"

func main() {
	score, valid := 5, false
	if score > 3 && valid {
		fmt.Println("good")
	} else {
		fmt.Println("not good")
	}
}
