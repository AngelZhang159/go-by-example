package main

import (
	"fmt"
	"time"
)

func goFunc(from string) {
	for i := 0; i < 300; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	//goFunc("direct")

	go goFunc("go1")
	go goFunc("go2")
	go goFunc("go3")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
