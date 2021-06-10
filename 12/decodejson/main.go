package main

type permissions map[string]bool

type user struct {
	user     string
	password string
	permissions
}

func main() {
	users := []user{
		{"inanc", "1234", nil},
		{"god", "42", map[string]bool{"admin": true}},
		{"devil", "666", map[string]bool{"write": true}},
	}
}
