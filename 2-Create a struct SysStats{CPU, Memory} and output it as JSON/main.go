// 2-Create a struct SysStats{CPU, Memory} and output it as JSON

package main

import (
	"encoding/json"
	"fmt"
)

// SysStats struct and JSON output
type SysStats struct {
	CPU    float64 `json:"cpu"`
	Memory uint64  `json:"memory"`
}

func printSysStatsJSON() {
	stats := SysStats{
		CPU:    42.7,
		Memory: 8096,
	}

	jsonData, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}

// Employee struct and methods
type Employee struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	IsRemote bool   `json:"isRemote"`
}

func (e *Employee) updateName(newName string) {
	e.Name = newName
}

func employeeDemo() {
	jobs := Employee{
		Name:     "sthita",
		Age:      12,
		IsRemote: true,
	}
	jobs.updateName("sthit")
	fmt.Println("employee name:", jobs.Name)
	fmt.Println("employee age:", jobs.Age)

	job := struct {
		title string
		id    int
	}{
		title: "devOps engineer",
		id:    345678,
	}
	fmt.Println("title role", job.title)
	fmt.Println("id num", job.id)
}

// Arrays demo
func arrayDemo() {
	arr1 := [5]int{1, 2, 4, 24, 4}
	sum := 0

	arr2D := [2][2]int{{1, 2}, {3, 2}}
	fmt.Println("arr2D[0][1]:", arr2D[0][1])

	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
	}

	fmt.Println("arr1:", arr1)
	fmt.Println("sum:", sum)
}

// Slices demo
func sliceDemo() {
	var z [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int = z[1:3]
	fmt.Println("len(s[:cap(s)]):", len(s[:cap(s)]))

	var a []int = []int{4, 45, 6, 6, 6, 54, 43, 3}
	a = append(a, 10) // add 10 at the end

	fmt.Printf("type of a: %T\n", a)
	fmt.Println("a:", a)
}

func mapDemo() {
	// Create a map with string keys and int values
	ages := map[string]int{
		"alice": 25,
		"bob":   30,
	}

	// Add a new key-value pair
	ages["carol"] = 22

	// Access a value
	fmt.Println("Alice's age:", ages["alice"])

	// Loop through the map
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Check if a key exists
	if age, ok := ages["dave"]; ok {
		fmt.Println("Dave's age:", age)
	} else {
		fmt.Println("Dave not found")
	}
}

// Main function calls all demos
func main() {
	printSysStatsJSON()
	employeeDemo()
	arrayDemo()
	sliceDemo()
	mapDemo()
}
