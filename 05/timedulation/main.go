package main

func main() {
	const (
		Nanosecond  Duration = 1
		Microsesond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsesond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Second
	)
}
