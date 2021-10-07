package main

import (
	"fmt"
	"time"
)

func inTimeSpan(start, end, check time.Time) bool {
	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
}

func main() {
	test := []struct {
		start string
		end   string
		check string
	}{
		{"08:00", "05:00", "12:00"},
	}
	newLayout := "12:00"
	for _, t := range test {
		check, _ := time.Parse(newLayout, t.check)
		start, _ := time.Parse(newLayout, t.start)
		end, _ := time.Parse(newLayout, t.end)
		fmt.Println(t.start+"-"+t.end, t.check, inTimeSpan(start, end, check))
	}
}
