package main

import (
	"fmt"
	"strings"
)

func main() {
	speed := 10
	fast := speed >= 80
	slow := speed < 20

	fmt.Printf("fast`s type is %T\n", fast)
	fmt.Printf("going fast? %t\n", fast)
	fmt.Printf("going slow? %t\n", slow)

	fmt.Printf("is it 100 mph? %t\n", speed == 100)
	fmt.Printf("is it not 100 mph? %t\n", speed != 100)

	fmt.Println(strings.Repeat("-", 25))

	speedB := 150.5
	faster := speedB > 100

	fmt.Println("is the other car going faster?", faster)
}
