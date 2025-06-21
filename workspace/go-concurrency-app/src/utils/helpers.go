package utils

import (
	//"fmt"
	"time"
)

func TaskOne() string {
	time.Sleep(2 * time.Second) // Simulate work
	return "Result from Task One"
}

func TaskTwo() string {
	time.Sleep(1 * time.Second) // Simulate work
	return "Result from Task Two"
}

func TaskThree() string {
	time.Sleep(3 * time.Second) // Simulate work
	return "Result from Task Three"
}