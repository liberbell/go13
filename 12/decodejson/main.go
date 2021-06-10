package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool

type user struct {
	User        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perms"`
}

func main() {
	users := []user{
		{"inanc", "1234", nil},
		{"god", "42", permissions{"admin": true}},
		{"devil", "666", permissions{"write": true}},
	}

	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}
