package main

import (
	"fmt"
	"os"
)

func main() {
	city := os.Args[1]

	switch city {
	case "tokyo":
		fmt.Println("Japan")
	case "paris":
		fmt.Println("France")
	}
}
