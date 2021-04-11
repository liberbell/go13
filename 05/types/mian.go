package main

func main() {
	type (
		Gram     int
		Kilogram Gram
		Ton      Kilogram
	)

	var (
		salt Gram = 100
		apple = Kilogram = 5
		truck Ton = 10
	)
}
