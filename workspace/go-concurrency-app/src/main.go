package main

import (
	"fmt"
	"time"
)

// Utility function that simulates a task
func task(name string, duration time.Duration, ch chan<- string) {
	time.Sleep(duration)
	ch <- fmt.Sprintf("Task %s completed", name)
}

func main() {
	// Create channels for communication
	task1Ch := make(chan string)
	task2Ch := make(chan string)
	task3Ch := make(chan string)

	// Start tasks concurrently
	go task("A", 2*time.Second, task1Ch)
	go task("B", 1*time.Second, task2Ch)
	go task("C", 3*time.Second, task3Ch)

	// Use select to wait for the first task to complete
	for i := 0; i < 3; i++ {
		select {
		case msg := <-task1Ch:
			fmt.Println(msg)
		case msg := <-task2Ch:
			fmt.Println(msg)
		case msg := <-task3Ch:
			fmt.Println(msg)
		}
	}

	fmt.Println("All tasks completed.")
}