package main

import (
	"fmt"
	"unsafe"

	s "github.com/inancgumus/prettyslice"
)

type collection []string

func main() {
	s.PrintElementAddr = true

	data := collection{"slices", "are", "awesome", "period"}
	fmt.Printf("slice`s size: %d bytes\n", unsafe.Sizeof(data))
	change(data)
	s.Show("main`s data ", data)
	fmt.Printf("main`s data slice address: %p\n", &data)

	array := [...]string{"slices", "are", "awesome", "period"}
	fmt.Printf("array`s size: %d bytes\n", unsafe.Sizeof(array))
	fmt.Printf("slice`s size: %d bytes\n", unsafe.Sizeof(data))
}

func change(data collection) {
	// var data collection
	// local data =
	data[2] = "brilliant"

	s.Show("change`s data: ", data)
	fmt.Printf("change`s data slice address: %p\n", &data)
}
