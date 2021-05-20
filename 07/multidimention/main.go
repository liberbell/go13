package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// spendings := [][]int{
	// 	{200, 100},
	// 	{500},
	// 	{50, 25, 75},
	// }
	// spendings := make([][]int, 0, 5)
	// spendings = append(spendings, []int{200, 100})
	// spendings = append(spendings, []int{35, 15, 45, 60})
	// spendings = append(spendings, []int{5, 15, 40})
	// spendings = append(spendings, []int{95, 10}, []int{50, 25})
	// spendings = append(spendings, []int{20, 130})

	spendings := fetch()

	for i, daily := range spendings {
		// fmt.Println(i+1, daily)
		var total int
		for _, spendings := range daily {
			total += spendings
		}

		fmt.Printf("Day %d: %d\n", i+1, total)
	}
}

func fetch() [][]int {
	contents := `200 100
35 15 45 60
5 15 40
95 10
50 25`

	lines := strings.Split(contents, "\n")
	spendings := make([][]int, len(lines))
	// fmt.Printf("%#v\n", spendings)

	for i, line := range lines {
		fmt.Printf("%d: %#v\n", i+1, line)

		fields := strings.Fields(line)

		spendings[i] = make([]int, len(fields))

		for j, field := range fields {
			fmt.Printf("\t%d: %#v\n", j+1, field)

			spending, _ := strconv.Atoi(field)
			spendings[i][j] = spending
		}
	}

	return spendings
}
