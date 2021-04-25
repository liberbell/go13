package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give a month name.")
		// return
	}
	m := os.Args[1]

	if m == "Dec" || m == "Jan" || m == "Feb" {
		fmt.Println("Winter")
	}
}
