package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	go func() {
		queue <- "one"
		queue <- "two"
		close(queue)
	}()

	//only use range if the channel is closed, or it will deadlock
	for s2 := range queue {
		fmt.Println("item:", s2)
	}

}
