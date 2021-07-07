package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	l := list{
		{title: "mody dick", price: 10, released: toTimestamp(118281600)},
		{title: "odyssey", price: 10, released: toTimestamp("723622855")},
		{title: "hobbit", price: 25},
	}

	data, err := json.MarshalIndent(l, "", " ")
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
