package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 7)
	a := 11
	go worker(c, 1)
	go worker(c, 2)
	go worker(c, 3)
	go generateInts(c, a)

	time.Sleep(1 * time.Second)
}

func worker(c chan int, id int) {
	for i := range c {
		fmt.Printf("Worker %d received value %d\n", id, i)
		time.Sleep(1 * time.Second)
		fmt.Println("Executing Goroutines")
	}
}

func generateInts(c chan int, a int) {
	for i := 0; i < a; i++ {
		c <- i
		fmt.Printf("Sent value %d\n", i)
		time.Sleep(1 * time.Millisecond)
	}
	close(c)
}

