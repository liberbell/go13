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
		fmt.Println("Good night.")
	case nowHour >= 12:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good morning.")
	}
}
