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
	m := os.Args[1]

	if m == "Dec" || m == "Jan" || m == "Feb" {
		fmt.Println("Winter")
	} else if m == "Mar" || m == "Apr" || m == "May" {
		fmt.Println("Spring")
	} else if m == "Jun" || m == "Jly" || m == "Aug" {
		fmt.Println("Summer")
	} else if m == "Sep" || m == "Oct" || m == "Nov" {
		fmt.Println("Autumn")
	} else {
		fmt.Printf("%q is wrong month.\n", m)
	}

	// switch m := os.Args[1]; m {
	// case "Dec", "Jan", "Feb":
	// 	fmt.Println("winter")
	// case "Mar", "Apr", "May":
	// 	fmt.Println("Spring")
	// case "Jun", "Jly", "Aug":
	// 	fmt.Println("Summer")
	// case "Sep", "Oct", "Nov":
	// 	fmt.Println("Autumn")
	// default:
	// 	fmt.Printf("%q is wrong month.\n", m)
	// }
}
