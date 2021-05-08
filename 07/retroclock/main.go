package main

import "fmt"

func main() {
	type placeholder [5]string

	zero := placeholder{
		"■■■",
		"■ ■",
		"■ ■",
		"■ ■",
		"■■■",
	}

	one := placeholder{
		"■■ ",
		" ■ ",
		" ■ ",
		" ■ ",
		"■■■",
	}

	two := placeholder{
		"■■■",
		"  ■",
		"■■■",
		"■  ",
		"■■■",
	}

	three := placeholder{
		"■■■",
		"  ■",
		"■■■",
		"  ■",
		"■■■",
	}

	four := placeholder{
		"■ ■",
		"■ ■",
		"■■■",
		"  ■",
		"  ■",
	}

	five := placeholder{
		"▉▉▉",
		"▉  ",
		"▉▉▉",
		"  ▉",
		"▉▉▉",
	}

	six := placeholder{
		"▉▉▉",
		"▉  ",
		"▉▉▉",
		"▉ ▉",
		"▉▉▉",
	}

	seven := placeholder{
		"▉▉▉",
		"  ▉",
		"  ▉",
		"  ▉",
		"  ▉",
	}

	eight := placeholder{
		"▉▉▉",
		"▉ ▉",
		"▉▉▉",
		"▉ ▉",
		"▉▉▉",
	}

	nine := placeholder{
		"▉▉▉",
		"▉ ▉",
		"▉▉▉",
		"  ▉",
		"▉▉▉",
	}

	// fmt.Println(zero, one, two, three, four, five, six, seven, eight, nine)

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	// fmt.Println(digits[0])
	// for _, digit := range digits {
	// 	for _, line := range digit {
	// 		fmt.Println(line)
	// 	}
	// }
	for line := range digits[0] {
		for digit := range digits {
			fmt.Print(digits[digit][line], "  ")
		}
		fmt.Println()
	}
}
