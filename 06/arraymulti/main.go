package main

import "fmt"

func main() {
	// students1 := [3]float64{5, 6, 1}
	// students2 := [3]float64{9, 8, 4}

	// var sum float64
	// sum += students1[0] + students1[1] + students1[2]
	// sum += students2[0] + students2[1] + students2[2]

	// const N = float64(len(students1) * 2)
	// fmt.Printf("Avg Grade: %g\n", sum/N)
	students := [...][3]float64{
		{5, 6, 1},
		{9, 8, 4},
	}
	var sum float64
	for _, grade := range students {
		for _, grade := range grade {
			sum += grade
		}
	}
	sum += students[0][0] + students[0][1] + students[0][2]
	sum += students[1][0] + students[1][1] + students[1][2]

	const N = float64(len(students) * len(students[0]))
	fmt.Printf("Avg Grade: %g\n", sum/N)
}
