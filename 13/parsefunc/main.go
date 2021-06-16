package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	p := newParser()

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p.lines++

		parsed, err := parse(p, in.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		update(p, parsed)

	}
}
