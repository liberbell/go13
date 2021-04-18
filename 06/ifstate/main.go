package main

import "fmt"

func main() {
	// score, valid := 5, false
	// if score > 3 && valid {
	// 	fmt.Println("good")
	// } else {
	// 	fmt.Println("not good")
	// }
	score := 3

	if score > 3 {
		fmt.Println("good")
	} else if score == 3 {
		fmt.Println("")
	}
}
