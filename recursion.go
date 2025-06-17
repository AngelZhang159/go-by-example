package main

import "fmt"

func main() {

	fmt.Println(fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(50))
}

func fact(i int) int {
	if i == 0 {
		return 1
	}
	return i * fact(i-1)
}
