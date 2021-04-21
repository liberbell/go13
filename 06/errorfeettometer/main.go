package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSapce(usage))
		return
	}
}
