package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}
	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		// fmt.Println(file.Name(), file.Size())
		if file.Size() == 0 {
			name := file.Name()
			fmt.Println(name)
		}
	}
}
