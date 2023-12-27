package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker(id int, jobs <-chan int, res chan int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		res <- j
	}
}

func main() {

	var numJobs = runtime.NumCPU()
	runtime.GOMAXPROCS(numJobs)

	jobs := make(chan int, numJobs)
	res := make(chan int, numJobs)

	for w := 1; w <= numJobs; w++ {
		go worker(w, jobs, res) // start 24 worker as machine cpu cores count
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		fmt.Println(<-res) // wait finish jobs
	}

}
