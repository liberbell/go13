package main

import "fmt"

type song struct {
	title, artist string
}

func main() {
	song1 := song{title: "wonderwall", artist: "Oasis"}
	song2 := song{title: "super sonic", artist: "Oasis"}

	fmt.Printf("Song1: %+v\n", song1)
	fmt.Printf("Song2: %+v\n", song2)

	if song1.title == song2.title && song1.artist == song2.artist {
		fmt.Println("song are equal.")
	}
}
