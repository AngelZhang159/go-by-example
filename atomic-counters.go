package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	//create atomic int that is safe to use with goroutines
	//it is another way of interacting with data across goroutines
	var ops atomic.Uint64

	//waitgroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	//create 50 goroutines, each adds 1000 to ops
	for range 50 {
		wg.Add(1)
		go func() {
			for range 1000 {
				ops.Add(1)
			}

			wg.Done()
		}()
	}

	//wait for all to finish
	wg.Wait()

	//prints 1000 * 50
	fmt.Println("ops", ops.Load())

	//version with normal int -> different result each time

	//var ops int64
	//
	//var wg sync.WaitGroup
	//
	//for range 50 {
	//	wg.Add(1)
	//	go func() {
	//		for range 1000 {
	//			ops += 1
	//		}
	//
	//		wg.Done()
	//	}()
	//}
	//
	//wg.Wait()
	//
	//fmt.Println("ops", ops)
}
