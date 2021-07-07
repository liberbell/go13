package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	l := list{
		{Title: "mody dick", Price: 10, Released: toTimestamp(118281600)},
		{Title: "odyssey", Price: 10, Released: toTimestamp("723622855")},
		{Title: "hobbit", Price: 25},
	}

	data, err := json.MarshalIndent(l, "", " ")
	if err != nil {
		log.Fatal(err)
		// fmt.Println(err)
		// return
	}

	fmt.Println(string(data))
}
