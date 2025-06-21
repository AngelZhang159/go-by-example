package main

import "fmt"

func main() {
	pings := make(chan string)
	pongs := make(chan string)

	go ping(pings, "passed message")
	go pong(pings, pongs)

	fmt.Println(<-pongs)
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}
