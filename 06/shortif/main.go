package main

import (
	"fmt"
	"strconv"
)

func main() {

	if n, err := strconv.Atoi("42"); err == nil {
		fmt.Println("There was no err: n is", n)
	}
}
