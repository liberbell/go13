package main

import "fmt"

func main() {
	word := []byte("力は5V/3A")
	fmt.Printf("%s = % [1]x\n", word)

	var size int
	for i := range string(word) {
		if i > 0 {
			size = 1
			break
		}
	}
}
