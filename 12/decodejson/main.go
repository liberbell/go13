package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool

type user struct {
	user     string
	password string
	permissions
}

func main() {
	users := []user{
		{"inanc", "1234", nil},
		{"god", "42", permissions{"admin": true}},
		{"devil", "666", permissions{"write": true}},
	}

	out, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}
}
