package main

import "fmt"

func main() {

	//buffered channel of 1 item
	message := make(chan string, 1)

	//if this channel was not buffered and the message was sent from the same goroutine as the receiver, it would deadlock because it is expecting a receiver

	//because it is buffered, it will continue the program after sending one item because it gets saved in the buffer and later processed by the receiving message in the same goroutine

	message <- "hello"
	//go func() { message <- "hello" }()

	fmt.Println(<-message)
}
