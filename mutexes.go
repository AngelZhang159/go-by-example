package main

import (
	"fmt"
	"sync"
)

// Container with a mutex so that it can be manipulated safely by multiple goroutines by locking and unlocking it when used
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	//locking and defer unlock (make it unlock when the func returns) then manipulate data safely
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {

	//the struct with initial data
	cont := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	fmt.Println(cont.counters)

	//waitgroup so that it waits for all the func to finish
	var wg sync.WaitGroup

	wg.Add(3)

	//anon func assigned to var / function literal / function value
	doIncrement := func(name string, n int) {
		for range n {
			cont.inc(name)
		}

		wg.Done()
	}

	//calls it safely in different goroutines, they will execute properly because of the waitgroup, without it, the main goroutine would finish before these even started
	go doIncrement("a", 50000)
	go doIncrement("a", 50000)
	go doIncrement("b", 50000)

	wg.Wait()
	fmt.Println(cont.counters)
}
