package main

import (
	"strings"

	"github.com/inancgumus/prettyslice"
)

func main() {
	// s := make([]int, 0, 5)
	// fmt.Println(s)
	// s = append(s, 42)
	// fmt.Println(s)

	prettyslice.PrintBacking = true
	prettyslice.MaxPerLine = 10

	tasks := []string{"jump", "run", "read"}
	// var upTasks []string

	upTasks := make([]string, 0, len(tasks))
	upTasks = upTasks[:cap(upTasks)]
	prettyslice.Show("upTasks", upTasks)

	for i, task := range tasks {
		upTasks[i] = strings.ToUpper(task)
		prettyslice.Show("upTasks", upTasks)
	}
}
