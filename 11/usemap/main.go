package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please type one search word.")
		return
	}
	query := args[0]
	// rx := regexp.MustCompile(`[^a-z]+`)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words := make(map[string]bool)

	for in.Scan() {
		// fmt.Println(in.Text())
		word := strings.ToLower(in.Text())
		// word = rx.ReplaceAllString(word, "")

		if len(word) > 2 {
			words[word] = true
		}
	}
	// fmt.Println("sun:", words["sun"], "tesla:", words["tesla"])

	// query := "many"
	if words[query] {
		fmt.Printf("The input contains %q\n", query)
	}

	fmt.Printf("Sorry, input does not contain %q.\n", query)
}
