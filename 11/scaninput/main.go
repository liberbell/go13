package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()

	fmt.Println("Sccaned text: ", in.Text())
}
