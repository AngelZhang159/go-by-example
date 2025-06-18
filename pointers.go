package main

import "fmt"

func main() {

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("value:", i)

	zeroptr(&i)
	fmt.Println("pointer:", i)
	fmt.Println("pointer value", &i)
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func zeroval(ival int) {
	ival = 0
}
