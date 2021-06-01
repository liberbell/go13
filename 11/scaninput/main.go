package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	os.Stdin.Close()
	in := bufio.NewScanner(os.Stdin)

	// for in.Scan() {
	// 	fmt.Println("Scanned text: ", in.Text())
	// }
	// fmt.Println("Sccaned text: ", in.Text())
	// fmt.Println("Sccaned text: ", in.Bytes())

	var lines int

	for in.Scan() {
		lines++
	}
	fmt.Printf("There are %d line(s)\n", lines)

	if err := in.Err(); err != nil {
		fmt.Println("Error: ", err)
	}
}
