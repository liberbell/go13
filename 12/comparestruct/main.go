package main

import "fmt"

func main() {
	song1 := song{title: "wonderwall", artist: "Oasis"}
	song2 := song{title: "super sonic", artist: "oasis"}

	fmt.Printf("Song1: %+v\n", song1)
	fmt.Printf("Song2: %+v\n", song2)
}
