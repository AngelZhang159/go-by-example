package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c <- "message"
	}()

	select {
	case val := <-c:
		fmt.Println("received 1:", val)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout first call")
	}

	go func() {
		time.Sleep(2 * time.Second)
		c <- "message"
	}()

	select {
	case val := <-c:
		fmt.Println("received 2:", val)
	case <-time.After(time.Second):
		fmt.Println("timeout second call")
	}
}
