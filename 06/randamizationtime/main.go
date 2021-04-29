package main

import (
	"fmt"
	"time"
)

func main() {
	unixt := time.Now()
	fmt.Println(unixt)

	unixsecond := unixt.Nanosecond()
	fmt.Println(unixsecond)
}
