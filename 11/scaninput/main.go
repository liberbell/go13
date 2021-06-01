package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		fmt.Println("Scanned text: ", in.Text())
	}

	// fmt.Println("Sccaned text: ", in.Text())
	// fmt.Println("Sccaned text: ", in.Bytes())
}
