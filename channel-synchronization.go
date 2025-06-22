package main

import (
	"fmt"
	"time"
)

func main() {
	//go terminate the program if it reaches the of the main func

	done := make(chan bool, 1)

	//go worker(done)
	go worker(done)

	//even if the go workers are not finished, if these channels do not wait for a response of the worker, the program will end

	//with this, the program will block until it receives a message from the channel
	<-done

	//if you don't have any calls to the goroutine but still expect a message, it will deadlock because nobody will ever send it, to fix this you do this:

	//select {
	//case val := <-done:
	//	fmt.Println("do something", val)
	//case <-time.After(2 * time.Second):
	//	fmt.Println("TIMEOUT: no message")
	//}

}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}
