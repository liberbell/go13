package main

import (
	"fmt"
	"os"
)

func main() {
	city := os.Args[1]

	switch city {
	case "Tokyo", "Osaka":
		fmt.Println("Japan")
	case "Paris", "Lyon":
		fmt.Println("France")
	default:
		fmt.Println("Other country.")
	}
}
