package main

import ("fmt", "unsafe")

func main() {
	empty = ""
	dump(empty)

}

type StringHeader struct {
	pointer uintptr
	length  int
}

func dump(s string) {
	ptr := *(*StringHeader)(unsafe.pointer(&s))
	fmt.Printf("%q: %+v\n", s, ptr)
}
