package main

func main() {
	str := "あいうえお 🙊"

	bytes := []byte(str)
	bytes[0] = 'N'
	bytes[1] = 'O'
}
