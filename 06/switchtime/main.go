package main

import (
	"fmt"
	"time"
)

func main() {
	nowHour := time.Now().Hour()
	fmt.Println("Current hour is: ", nowHour)

	switch {
	case nowHour >= 18:
		fmt.Println("Good evening.")
	case nowHour >= 12:
		fmt.Println("Good afternoon.")
	case nowHour >= 6:
		fmt.Println("Good morning.")
	default:
		fmt.Println("Good night.")
	}
}
