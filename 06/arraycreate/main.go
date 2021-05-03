package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	books := [yearly]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
	}

	books[0] = "Kafka's Revenge"
	books[1] = "Stay Golden"
	books[2] = "Everythingship"
	books[3] += books[0] + " 2nd Edition"

	// fmt.Printf("books: %T\n", books)
	// fmt.Println("books: ", books)
	// fmt.Printf("books: %q\n", books)
	fmt.Printf("books: %#v\n", books)

	// var (
	// 	wBooks [winter]string
	// 	sBooks [summer]string
	// )

	// wBooks[0] = books[0]
	// sBooks[0] = books[1]
	// sBooks[1] = books[2]
	// sBooks[2] = books[3]
	// for i := range sBooks {
	// 	sBooks[i] = books[i+1]
	// }

	// for _, v := range sBooks {
	// 	// sBooks[i] = books[i+1]
	// 	v += "won`t effect"
	// 	// fmt.Println(v)
	// }

	// fmt.Printf("\nwinter : %#v\n", wBooks)
	// fmt.Printf("\nsummer : %#v\n", sBooks)

	// var published [len(books)]bool
	// published[0] = true
	// published[len(books)-1] = true

	// fmt.Println("\nPublished books:")
	// for i, ok := range published {
	// 	if ok {
	// 		fmt.Printf("+ %s\n", books[i])
	// 	}
	// }

}
