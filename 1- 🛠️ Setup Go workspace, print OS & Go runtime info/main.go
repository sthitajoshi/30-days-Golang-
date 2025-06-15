package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("OS:", os.Getenv("OS"))
}
