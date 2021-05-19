package main

import "fmt"

func main() {
	// spendings := [][]int{
	// 	{200, 100},
	// 	{500},
	// 	{50, 25, 75},
	// }
	spendings := make([][]int, 0, 5)
	spendings = append(spendings, []int{200, 100})
	spendings = append(spendings, []int{35, 15, 45, 60})
	spendings = append(spendings, []int{5, 15, 40})
	spendings = append(spendings, []int{95, 10}, []int{50, 25})

	for i, daily := range spendings {
		// fmt.Println(i+1, daily)
		var total int
		for _, spendings := range daily {
			total += spendings
		}

		fmt.Printf("Day %d: %d\n", i+1, total)
	}
}
