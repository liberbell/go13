package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give a month name.")
		return
	}
	// m := os.Args[1]

	// if m == "Dec" || m == "Jan" || m == "Feb" {
	// 	fmt.Println("Winter")
	// } else if m == "Mar" || m == "Apr" || m == "May" {
	// 	fmt.Println("Spring")
	// } else if m == "Jun" || m == "Jly" || m == "Aug" {
	// 	fmt.Println("Summer")
	// } else if m == "Sep" || m == "Oct" || m == "Nov" {
	// 	fmt.Println("Autumn")
	// } else {
	// 	fmt.Printf("%q is wrong month.\n", m)
	// }

	switch m := os.Args[1]; m {
	case m == "Dec" || m == "Jan" || m == "Feb":
		fmt.Println("winter")
	}
}
