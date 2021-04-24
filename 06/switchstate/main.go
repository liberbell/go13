package main

import "fmt"

func main() {
	city := "paris"

	switch city {
	case "tokyo":
		fmt.Println("Japan")
	case "paris":
		fmt.Println("France")
	}
}
