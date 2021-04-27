package main

import "fmt"

func main() {
	var (
		sum int
		i   = 1
	)

	for {
		if i > 5 {
			break
		}

		if i%2 != 0 {
			fmt.Println(i)
			i++
			continue
		}

		sum += i
		fmt.Println(i, "-->", sum)
		i++
	}
	fmt.Println(sum)
}
