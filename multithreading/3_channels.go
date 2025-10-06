package multithreading

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- string) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- fmt.Sprintf("worker %d result %d", id, j*2)
	}
}

func RunWorkersWithBuffering() {
	const (
		jobsCount    = 5
		workersCount = 3
	)
	jobs := make(chan int, jobsCount)
	results := make(chan string, jobsCount)
	defer close(results)

	for i := 1; i <= jobsCount; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 1; i <= workersCount; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= jobsCount; i++ {
		fmt.Println(<-results)
	}
}

func RunWorkersWithoutBuffering() {
	const (
		jobsCount    = 5
		workersCount = 3
	)
	jobs := make(chan int)
	results := make(chan string)
	defer close(results)

	for i := 1; i <= workersCount; i++ {
		go worker(i, jobs, results)
	}

	go func() {
		for i := 1; i <= jobsCount; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	for i := 1; i <= jobsCount; i++ {
		fmt.Println(<-results)
	}
}
