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
	// copied := map[string]string{"down": "した", "up": "うえ", "good": "グッド", "great": "すばらしい", "mistake": "", "perfect": "かんぺき"}

	// first := fmt.Sprintf("%s", dict)
	// second := fmt.Sprintf("%s", copied)

	// if first == second {
	// 	fmt.Println("Maps are equal.")
	// }

	// key := "good"

	// value, ok := dict[query]
	// if !ok {
	// 	fmt.Printf("%q not found\n", query)
	// 	return
	// }

	japanese := make(map[string]string, len(dict))

	for k, v := range dict {
		japanese[v] := k
	}
	
	japanese["good"] = "わるくない"
	dict["great"] = "さいこう"
	fmt.Printf("english: %q\njapanese: %q\n", dict, japanese)

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
	// 	fmt.Printf("%q means %q\n", w, japanese[i])
	// 		return
	// 	}
	// }
	fmt.Printf("%q not found\n", query)
}
