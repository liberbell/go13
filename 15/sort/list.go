package main

type list []*product

func (l list) String() string {}

func (l list) discount(ratio float64) {
	// abc
}

func (l list) len() int {
	return len(l)
}

func (l list) Less(i, j int) bool {
	return l[i].title < l[j].title
}

func (l list) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
