package main

import (
	"fmt"
	"os"
)

func main() {
	var name string

	name = os.Args[1]
	name, age := "gandalf", 2019

	fmt.Println("Hello great", name, "!")
	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("BTW")
}
