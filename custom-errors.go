package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func fun(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{
			arg:     arg,
			message: "can't work with it",
		}
	}

	return arg + 3, nil
}

func main() {

	_, err := fun(42)

	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
