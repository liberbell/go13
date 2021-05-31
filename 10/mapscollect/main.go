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
	dict := map[string]string{
		"good":    "よい",
		"great":   "すばらしい",
		"perfect": "かんぺき",
	}

	dict["good"] = "グッド"

	dict["up"] = "うえ"
	dict["down"] = "した"
	dict["mistake"] = ""

	// for k, v := range dict {
	// 	fmt.Printf("%q means %#v\n", k, v)
	// }
	copied := map[string]string{"down": "した", "up": "うえ", "good": "グッド", "great": "すばらしい", "mistake": "", "perfect": "かんぺき"}

	// key := "good"

	// value, ok := dict[query]
	// if !ok {
	// 	fmt.Printf("%q not found\n", query)
	// 	return
	// }

	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %v\n", query, value)
		return
	}

	// fmt.Printf("%q means %#v\n", query, value)

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
