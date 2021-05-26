package main

func main() {
	empty = ""
	dump(empty)

}

type StringHeader {
	pointer uintptr
	length int
}