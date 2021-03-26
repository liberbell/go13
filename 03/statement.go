package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(" World" + "!")
	fmt.Println(runtime.NumCPU() + 1)

	if 5 > 1 {
		fmt.Println("Bigger")
	}
}
