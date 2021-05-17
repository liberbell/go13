package main

import (
	"fmt"
	"strings"

	"github.com/inancgumus/prettyslice"
)

func main() {
	s := make([]int, 0, 5)
	fmt.Println(s)
	s = append(s, 42)
	fmt.Println(s)

	prettyslice.PrintBacking = true
	prettyslice.MaxPerLine = 10

	tasks := []string{"jump", "run", "read"}
	var upTasks []string
	prettyslice.Show("upTasks", upTasks)

	for _, task := range tasks {
		upTasks = append(upTasks, strings.ToUpper(task))
		prettyslice.Show("upTasks", upTasks)
	}
}
