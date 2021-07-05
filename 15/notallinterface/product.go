package main

import (
	"fmt"
	"strconv"
	"time"
)

type product struct {
	title    string
	price    money
	released timestamp
}

func (p *product) print() {
	fmt.Printf("%s: %s (%s)\n", p.title, p.price, p.released.string())
}

func (p *product) discount(ratio float64) {
	p.price *= money(1 - ratio)
}

func toTimestamp(v interface{}) (ts timestamp) {
	var t int

	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	}

	ts.Time = time.Unix(int64(t), 0)
	return ts
}
