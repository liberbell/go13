package main

func (l []*game) print() {
	for _, it := range l {
		it.print()
	}
}
