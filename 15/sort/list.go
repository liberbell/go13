package main

type list []*product

func (l list) String() string {
	// abc
}

func (l list) discount(ratio float64) {
	// abc
}

func (l list) len() int {
	return len(l)
}

func (l list) Less(i, j int) bool {
	return l[i].title < l[j].title
}
