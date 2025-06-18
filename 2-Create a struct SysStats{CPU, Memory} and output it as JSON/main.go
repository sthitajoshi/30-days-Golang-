// 2-Create a struct SysStats{CPU, Memory} and output it as JSON

package main

import (
	"encoding/json"
	"fmt"
)

type SysStats struct {
	CPU    float64 `json:"cpu"`
	Memory uint64  `json:"memory"`
}

func main() {
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

// STRUCT

// type Employee struct {
// 	Name     string `json:"name"`
// 	Age      int    `json:"age"`
// 	IsRemote bool   `json:"isRemote"`
// }

// func (e *Employee) updateName(newName string) {
// 	e.Name = newName
// }

// func main() {
// 	jobs := Employee{
// 		Name:     "sthita",
// 		Age:      12,
// 		IsRemote: true,
// 	}
// 	jobs.updateName("sthit")
// 	fmt.Println("employee name:", jobs.Name)
// 	fmt.Println("employee age:", jobs.Age)

// 	job := struct {
// 		title string
// 		id    int
// 	}{
// 		title: "devOps engineer",
// 		id:    345678,
// 	}
// 	fmt.Println("title role", job.title)
// 	fmt.Println("id num", job.id)

// }

// // collections
// // arrays
// // moved array code inside main function
// // slices and maps can be added here as needed

// func main() {
// 	arr1 := [5]int{1, 2, 4, 24, 4}
// 	sum := 0

// 	arr2D := [2][2]int{{1, 2}, {3, 2}}
// 	fmt.Print(arr2D[0][1])

// 	for i := 0; i < len(arr1); i++ {

// 		sum += arr1[i]
// 	}

// 	fmt.Println(arr1)
// }

// // slices

// func main() {
// 	var z [5]int = [5]int{1, 2, 3, 4, 5}
// 	var s []int = z[1:3]
// 	fmt.Println(len(s[:cap(s)]))

// 	var a []int = []int{4, 45, 6, 6, 6, 54, 43, 3}
// 	a = append(a, 10) //add 10 in the end of the a

// 	fmt.Printf("%T\n", a)
// }

//maps
