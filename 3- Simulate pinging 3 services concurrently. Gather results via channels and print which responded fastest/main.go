// 3- Simulate pinging 3 services concurrently. Gather results via channels and print which responded fastest
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pingService(name string, ch chan string) {
	// Simulate variable response time
	delay := time.Duration(rand.Intn(100)) * time.Millisecond
	time.Sleep(delay)
	ch <- name
}

// Goroutine example
func somefunc(num string) {
	fmt.Println("Goroutine:", num)
}

// Channel and select example
func channelSelectExample() {
	myChanel := make(chan string)
	anotherChanel := make(chan string)
	go func() {
		myChanel <- "await"
	}()
	go func() {
		anotherChanel <- "sync"
	}()
	time.Sleep(time.Second * 2)
	msg := <-myChanel
	select {
	case msgFromMychanel := <-myChanel:
		fmt.Println("From myChanel:", msgFromMychanel)
	case msgFromAnotherchanel := <-anotherChanel:
		fmt.Println("From anotherChanel:", msgFromAnotherchanel)
	}
	fmt.Println("First message from myChanel:", msg)
}

// Concurrent tasks with select
func concurrentTasks() {
	// Utility function that simulates a task
	task := func(name string, duration time.Duration, ch chan<- string) {
		time.Sleep(duration)
		ch <- fmt.Sprintf("Task %s completed", name)
	}

	task1Ch := make(chan string)
	task2Ch := make(chan string)
	task3Ch := make(chan string)

	go task("A", 2*time.Second, task1Ch)
	go task("B", 1*time.Second, task2Ch)
	go task("C", 3*time.Second, task3Ch)

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

func forSelectLoop() {

	charChanel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		charChanel <- s
	}

	close(charChanel)

	for result := range charChanel {
		fmt.Println(result)
	}
}

func main() {

	rand.Seed(time.Now().UnixNano()) // Seed randomness

	ch := make(chan string)

	// Launch 3 service pings concurrently
	go pingService("Service A", ch)
	go pingService("Service B", ch)
	go pingService("Service C", ch)

	// Receive the first response
	fastest := <-ch
	fmt.Println("Fastest response from:", fastest)

	go somefunc("1")
	go somefunc("2")
	go somefunc("3")
	time.Sleep(time.Second)

	channelSelectExample()

	concurrentTasks()

	forSelectLoop()
}
