package main

import (
	"fmt"
	"sync"
	"time"
)

func workerWaitGroup(id int) {
	fmt.Println("starting job", id)
	time.Sleep(time.Second)
	fmt.Println("finished job", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			//defer executes a function before the parent function returns, used for things like closing files or wait groups
			defer wg.Done()
			workerWaitGroup(i)
		}()
	}

	wg.Wait()

	fmt.Println("program finished")
}
