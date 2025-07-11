package main

import (
	"fmt"
	"time"
)

func workerJobs(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "started job", job)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", job)
		result <- job * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	//create 3 workers that accept jobs and send results
	for i := 1; i <= 3; i++ {
		go workerJobs(i, jobs, results)
	}

	//send jobs to the chan
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	//indicate that jobs chan isn't going to be used anymore so that it closes the goroutines
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		fmt.Println(<-results)
	}
}
