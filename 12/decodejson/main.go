package main

type user struct {
	user        string
	password    string
	permissions map[string]bool
}

func main() {
	users := []user{
		{"inanc", "1234", nil},
		{"god", "42", map[string]bool{"admin": true}},
	}
}
