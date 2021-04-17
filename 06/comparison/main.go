package main

import "fmt"

func main() {
	speed := 100
	fast := speed >= 80

	fmt.Printf("fast`s type is %T\n", fast)
	fmt.Printf("going fast? %t\n", fast)
}
