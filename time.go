package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println(now)

	then := time.Date(2029, 11, 19, 20, 50, 15, 651387237, time.UTC)

	fmt.Println(then.Year())
	fmt.Println(then.Month())
	fmt.Println(then.Day())
	fmt.Println(then.Hour())
	fmt.Println(then.Minute())
	fmt.Println(then.Second())
	fmt.Println(then.Nanosecond())
	fmt.Println(then.Location())

	fmt.Println(then.Weekday())

	fmt.Println(then.Before(now))
	fmt.Println(then.After(now))
	fmt.Println(then.Equal(now))

	diff := now.Sub(then)
	fmt.Println(diff)

	fmt.Println(diff.Hours())
	fmt.Println(diff.Minutes())
	fmt.Println(diff.Seconds())

	fmt.Println(then.Add(diff))
	fmt.Println(then.Add(-diff))
}
