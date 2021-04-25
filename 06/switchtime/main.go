package main

import (
	"fmt"
	"time"
)

func main() {
	nowHour := time.Now().Hour()
	fmt.Println("Current hour is: ", nowHour)

	switch {
	case nowHour >= 6, nowHour < 12:
		fmt.Println("Good morning.")
	case nowHour >= 12, nowHour < 17:
		fmt.Println("Good evening.")
	}
}
