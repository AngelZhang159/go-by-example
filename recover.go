package main

import (
	"fmt"
)

func mayPanic() {
	panic("some problem")
}

func main() {

	//recover from a panic so that it doesn't crash the program

	//recover in a "defer" will stop the panic in the function it was called and continue execution, in this example, in the main, it will stop at the panic and return because there is no other upper function

	//if the panic and recover were inside another func then it would continue in the main
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered. error: ", r)
		}
	}()

	mayPanic()

	fmt.Println("this wont print because of the panic")
}
