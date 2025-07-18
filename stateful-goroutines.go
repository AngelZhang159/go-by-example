package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key   int
	value int
	resp  chan bool
}

func main() {

	//here we create a goroutine that holds the state
	//we can create other goroutines that read and write to this state, and they communicate via chans

	//counters for each op
	var readOpsCounter uint64
	var writeOpsCounter uint64

	//main communication channels, where they will send messages to read and write
	var reads = make(chan readOp)
	var writes = make(chan writeOp)

	//stateful goroutine that holds a map int / int and receives the read and write ops
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			}
		}
	}()

	//create 100 reapOps that add to the counter
	for range 100 {
		go func() {
			for {
				op := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- op
				<-op.resp
				atomic.AddUint64(&readOpsCounter, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	//create 10 writeOps that add to the counter
	for range 10 {
		go func() {
			for {
				op := writeOp{
					key:   rand.Intn(5),
					value: rand.Intn(100),
					resp:  make(chan bool),
				}

				//send op
				writes <- op
				//block until response
				<-op.resp
				atomic.AddUint64(&writeOpsCounter, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	//wait to see how many ops we do in 1s (the read & write goroutines are infinite)
	time.Sleep(time.Second)

	fmt.Printf("result: reads: %d, %d\n", atomic.LoadUint64(&readOpsCounter), atomic.LoadUint64(&writeOpsCounter))
}
