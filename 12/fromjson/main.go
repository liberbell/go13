package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name        string `json:"username"`
	Permissions map[string]bool
}

func main() {
	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
		// input = append(input, '\n')
	}
	// fmt.Println(string(input))
	var users []user

	err := json.Unmarshal(input, &users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(users)
}
