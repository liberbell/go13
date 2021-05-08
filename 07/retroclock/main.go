package main

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
}
