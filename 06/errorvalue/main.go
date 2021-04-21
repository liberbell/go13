package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// s := strconv.Itoa(42)
	// fmt.Println(s)
	// fmt.Printf("%T\n", s)
	// n, err := strconv.Atoi(os.Args[1])
	// fmt.Println("Convert number: ", n)
	// fmt.Println("Returned Err:  ", err)
	age := os.Args[1]
	n, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Printf("Success: Converted %q", n)
	}
}
