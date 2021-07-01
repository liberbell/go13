package main

func main() {
	nums := []int{1, 2, 3}

	var any interface{}
	any = nums

	_ = len(any)
}
