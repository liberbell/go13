package main

import "time"

type timestamp struct {
	time.Time
}

func (ts timestamp) string() string {
	if ts.IsZero() {
		return "unknown"
	}
}
