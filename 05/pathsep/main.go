package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string

	dir, file = path.Split("main.go")
	fmt.Println(dir, file)
}
