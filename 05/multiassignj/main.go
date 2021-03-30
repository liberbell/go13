package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		speed     = 100
		prevSpeed = 50
		now       time.Time
	)

	speed, now = 100, time.Now()
	fmt.Println(speed, now)

	fmt.Println(speed, prevSpeed)
	speed, prevSpeed = prevSpeed, speed
	fmt.Println(speed, prevSpeed)
}
