package main

import (
	"fmt"
)

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type element[T interface{}] struct {
	next *element[T]
	val  T
}

type List[T interface{}] struct {
	head *element[T]
	tail *element[T]
}

func (lst *List[T]) Push(v T) {
	if lst.head == nil {
		lst.head = &element[T]{
			val: v,
		}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{
			val: v,
		}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for i := lst.head; i != nil; i = i.next {
		elems = append(elems, i.val)
	}
	return elems
}

func main() {

	s := []string{"foo", "bar", "zoo"}

	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))
	// 	you can explicitly type the parameters [slice, string]
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(20)
	fmt.Println("list:", lst.AllElements())
}
