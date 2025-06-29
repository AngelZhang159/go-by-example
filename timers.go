package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)

	//this works like time.Sleep()
	fmt.Println("time start")
	<-timer1.C
	fmt.Println("time end")

	//the good thing about time.NewTimer is that it can be canceled before being fired
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("firing timer 2")
	}()

	//if the timer had been fired before reaching the stop line, the stop would fail and return false
	//time.Sleep(2 * time.Second)

	//cancelling timer
	stopped := timer2.Stop()

	if stopped {
		fmt.Println("timer 2 stopped")
	} else {
		fmt.Println("timer 2 fired")
	}

	time.Sleep(2 * time.Second)
}
