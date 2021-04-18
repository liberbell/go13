package main

import "fmt"

func main() {
	color := "red"
	fmt.Println("raddish color?", color == "red" || color == "dark red")

	color = "dark red"
	fmt.Println("raddish color?", color == "red" || color == "dark red")

	fmt.Println("greenish color?", color == "green" || color == "light green")
}
