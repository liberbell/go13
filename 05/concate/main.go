package main

import "fmt"

func main() {
	name, last := "carl", "sagan"
	name += " edward"
	fmt.Println(name + " " + last)

	fmt.Println(
		"hello" + ", " + "how" + " " + "are you" + " " +
			"today?",
	)

	fmt.Println(
		`hello` + `, ` + `how` + ` ` + `are you` + ` ` +
			`today?`,
	)
}
