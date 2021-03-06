package main

import (
	"fmt"
	"strconv"
)

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

	eq := "1 + 2 = "
	sum := 1 + 2
	fmt.Println(eq + strconv.Itoa(sum))

	// eq := true + " " + false
	eq = strconv.FormatBool(true) + " " + strconv.FormatBool(false)
	fmt.Println(eq)
}
