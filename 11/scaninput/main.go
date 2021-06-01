package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	in.Scan()
	in.Scan()

	fmt.Println("Sccaned text: ", in.Text())
	fmt.Println("Sccaned text: ", in.Text())
	fmt.Println("Sccaned text: ", in.Text())
}
