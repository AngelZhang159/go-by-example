package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	default:
		fmt.Println("no messages received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("message sent")
	default:
		fmt.Println("no receiver available")
	}

	//go func() {
	//	messages <- msg
	//}()

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
