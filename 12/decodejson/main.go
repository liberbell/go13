package main

type user struct {
	user string
	password string
	permissions map[string]bool
}

func main()  {
	users := []user{
		{"inanc": "123"}
	}
}