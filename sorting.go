package main

import (
	"fmt"
	"slices"
)

func main() {

	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("strings: ", strs)

	ints := []int{1, 5, 6, 11, 15}
	slices.Sort(ints)
	fmt.Println("ints: ", ints)

	s := slices.IsSorted(ints)
	fmt.Println("sorted?: ", s)
}
