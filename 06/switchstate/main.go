package main

import "fmt"

func main() {
	city := "tokyo"

	switch city {
	case "tokyo":
		fmt.Println("Japan")
	case "paris":
		fmt.Println("France")
	}
}
