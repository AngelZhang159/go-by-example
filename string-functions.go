package main

import (
	"fmt"
	. "strings"
)

func main() {

	fmt.Println("Contains:  ", Contains("test", "es"))
	fmt.Println("Count:     ", Count("test", "t"))
	fmt.Println("HasPrefix: ", HasPrefix("test", "te"))
	fmt.Println("HasSuffix: ", HasSuffix("test", "st"))
	fmt.Println("Index:     ", Index("test", "e"))
	fmt.Println("Join:      ", Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat:    ", Repeat("a", 5))
	fmt.Println("Replace:   ", Replace("foo", "o", "0", -1))
	fmt.Println("Replace:   ", Replace("foo", "o", "0", 1))
	fmt.Println("Split:     ", Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:   ", ToLower("TEST"))
	fmt.Println("ToUpper:   ", ToUpper("test"))

}
