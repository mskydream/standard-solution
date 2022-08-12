package main

import (
	"fmt"
	"time"
)

func main() {
	aboutTime()
	duration()
}

func aboutTime() {
	fmt.Printf("now: %v\n", time.Now())

	t := time.Date(2017, time.August, 1, 18, 25, 45, 0, time.UTC)
	fmt.Printf("past: %v\n", t)
	fmt.Printf("weekday: %s\n", t.Weekday())
	fmt.Printf("month: %d\n", t.Month())
	fmt.Printf("day: %d\n", t.Day())
	fmt.Printf("year: %d\n", t.Year())
	fmt.Printf("hour: %d\n", t.Hour())
	fmt.Printf("minute: %d\n", t.Minute())
	fmt.Printf("second: %d\n", t.Second())
	fmt.Printf("nanosecond: %d\n", t.Nanosecond())
	fmt.Printf("location: %s\n", t.Location())
}

func duration() {
	t1 := time.Now()
	fmt.Printf("t1: %s\n", t1)

	d := 1 * time.Hour
	t2 := t1.Add(d)
	fmt.Printf("t2: %s\n", t2)
	fmt.Printf("t2.Before(t1): %t\n", t2.Before(t1))
	fmt.Printf("t2.After(t1): %t\n", t2.After(t1))
	fmt.Printf("t2.Equal(t1): %t\n", t2.Equal(t1))

}
