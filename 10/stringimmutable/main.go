package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// empty = ""
	// dump(empty)
	hello := "hello"
	dump(hello)
	dump("hello")
	dump("hello!")

	dump(hello[0:1])

}

type StringHeader struct {
	pointer uintptr
	length  int
}

func dump(s string) {
	ptr := *(*StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%q: %+v\n", s, ptr)
}