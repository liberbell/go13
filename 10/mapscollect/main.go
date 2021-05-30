package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [japanese word]")
		return
	}

	query := args[0]

	english := []string{"good", "great", "perfect"}
	japanese := []string{"よい", "すげえ", "完璧"}

	for i, w := range english {
		if query == w {
			fmt.Printf("%q means %q\n", w, japanese[i])
			return
		}
	}
}
