package main

import (
	"fmt"
	"strconv"
	"time"
)

type book struct {
	title     string
	price     money
	published interface{}
}

func (b book) print() {
	p := format(b.published)
	fmt.Printf("%-15s: %s - (%v)\n", b.title, b.price.string(), p)
}

func format(v interface{}) string {
	var t int
	switch v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	default:
		return "unknown"
	}

	if v, ok := v.(string); ok {
		t, _ = strconv.Atoi(v)
	}

	u := time.Unix(int64(t), 0)
	return u.String()
}
