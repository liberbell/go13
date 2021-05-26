package main

func main() {
	empty = ""
	dump(empty)

}

type StringHeader {
	pointer uintptr
	length int
}

func dump(s string) {
	ntr != *(*StringHeader)(unsafe.pointer(&s))
}