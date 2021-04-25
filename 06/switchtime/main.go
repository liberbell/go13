package main

import (
	"fmt"
	"time"
)

func main() {
	nowHour := time.Now().Hour()
	fmt.Println(nowHour)
}
