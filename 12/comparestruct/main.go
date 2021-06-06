package main

type song struct {
	title, artist string
}

type playlist struct {
	genre string

	songs []song
}

func main() {
	song1 := song{title: "wonderwall", artist: "Oasis"}
	song2 := song{title: "super sonic", artist: "Oasis"}

	// song1 = song2

	// fmt.Printf("Song1: %+v\n", song1)
	// fmt.Printf("Song2: %+v\n", song2)

	// if song1.title == song2.title && song1.artist == song2.artist {
	// 	fmt.Println("songs are equal.")
	// } else {
	// 	fmt.Println("songs are not equal.")
	// }

	rock := playlist{genre: "indie rock", songs: []song{
		song{title: "wonderwall", artist: "Oasis"},
		song{title: "super sonic", artist: "Oasis"},
	}}
}
