package main

import (
	"fmt"
)

func main() {
	// args := os.Args[1:]
	// if len(args) != 1 {
	// 	fmt.Println("[english word] -> [japanese word]")
	// 	return
	// }

	// query := args[0]
	var dict map[string]string

	key := "good"
	if dict != nil {
		value := dict[key]
		fmt.Printf("%q means %#v\n", key, value)
	}

	// fmt.Printf("Zero value: %v\n", dict)
	fmt.Printf("# of keys : %d\n", len(dict))

	// english := []string{"good", "great", "perfect"}
	// japanese := []string{"よい", "すげえ", "完璧"}

	// for i, w := range english {
	// 	if query == w {
	// 		fmt.Printf("%q means %q\n", w, japanese[i])
	// 		return
	// 	}
	// }
	// fmt.Printf("%q not found\n", query)
}
