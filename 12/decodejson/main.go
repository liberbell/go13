package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool

type user struct {
	User        string
	Password    string
	Permissions permissions
}

func main() {
	users := []user{
		{"inanc", "1234", nil},
		{"god", "42", permissions{"admin": true}},
		{"devil", "666", permissions{"write": true}},
	}

	out, err := json.MarshalIndent(users, ">>", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}
