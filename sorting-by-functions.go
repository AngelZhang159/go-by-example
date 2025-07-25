package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {

	//sort by custom function

	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{
			name: "Paco",
			age:  2,
		},
		{
			name: "Laura",
			age:  5,
		},
		{
			name: "Jose",
			age:  188,
		},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people)

}
