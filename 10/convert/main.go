package main

import "fmt"

func main() {
	str := "あいうえお 🙉"

	bytes := []byte(str)
	bytes[0] = 'N'
	bytes[1] = 'O'

	str = string(bytes)

	fmt.Printf("%s\n", str)
}
