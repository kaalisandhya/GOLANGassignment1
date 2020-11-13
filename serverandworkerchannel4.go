1.package main

import (
	"fmt"
	"time"
)
func employee(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("employee", id, "started doing my job", j)
		time.Sleep(time.Second)
		fmt.Println("employee", id, "finished doing my job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
	go employee(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
$go run main.go
employee 3 started doing my job 1
employee 1 started doing my job 2
employee 2 started doing my job 3
employee 3 finished doing my job 1
employee 3 started doing my job 4
employee 1 finished doing my job 2
employee 2 finished doing my job 3
employee 1 started doing my job 5
employee 3 finished doing my job 4
employee 1 finished doing my job 5
package main
import (
    "fmt"
    "time"
)
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
       results <- j * 2
    }
}
func main() {
  const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
for j := 1; j <= numJobs; j++ {
        jobs <- j
}
    close(jobs)
for a := 1; a <= numJobs; a++ {
        <-results
}
}
$go run main.go
worker 3 started  job 1
worker 1 started  job 2
worker 2 started  job 3
worker 3 finished job 1
worker 3 started  job 4
worker 1 finished job 2
worker 1 started  job 5
worker 2 finished job 3
worker 3 finished job 4
worker 1 finished job 5
