package main

import "fmt"

func main() {
	speed := 100
	fast := speed >= 80
	slow := speed < 20

	fmt.Printf("fast`s type is %T\n", fast)
	fmt.Printf("going fast? %t\n", fast)
	fmt.Printf("going slow? %t\n", fast)
}
