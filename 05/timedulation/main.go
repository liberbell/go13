package main

import (
	"fmt"
	"time"
)

func main() {
	// const (
	// 	Nanosecond  Duration = 1
	// 	Microsesond          = 1000 * Nanosecond
	// 	Millisecond          = 1000 * Microsesond
	// 	Second               = 1000 * Millisecond
	// 	Minute               = 60 * Second
	// 	Hour                 = 60 * Second
	// )

	// t := time.Second * 10
	// fmt.Println(t)

	// i := 10
	// t = time.Second * time.Duration(i)
	// fmt.Println(t)
	const c = 10
	t := time.Second * c
	fmt.Println(t)

}
