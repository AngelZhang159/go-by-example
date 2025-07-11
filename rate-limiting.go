package main

import (
	"fmt"
	"time"
)

func main() {
	//make a requests chan with buffer and send 5 messages then close
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	//make time ticker chan of 200ms
	limiter := time.Tick(200 * time.Millisecond)

	for request := range requests {
		//receiving the limiter chan block the goroutine for 200ms
		<-limiter
		//result: print every 200ms 5 times
		fmt.Println("request", request, time.Now())
	}

	//-------------------------------------

	//create limiter chan with buffer 3, it holds time data
	burstyLimiter := make(chan time.Time, 3)

	//send 3 messages to fill buffer first
	for range 3 {
		burstyLimiter <- time.Now()
	}

	//later, send a message every 200ms
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	//create chan to send messages and send 5 instantly
	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	//receive the request data
	for req := range burstyRequests {
		//block goroutine to the buffer size of the limiter (3)
		//it will only be able to receive data every 200ms, and only the size of the burstyLimiter's buffer, in this case, it has 3 first, then gets one each 200ms, it will only be able to process req when the limiter gets a message
		<-burstyLimiter
		fmt.Println("requests", req, time.Now())
	}
}
