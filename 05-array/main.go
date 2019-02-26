package main

import "fmt"

func main() {
	var arr [5]int

	b := arr
	b[0] = 100
	b[2] = 200
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	// a := [...]string{"a", "b", "c", "d"}

	fmt.Println(b[0])

	fmt.Println()
	fmt.Println(arr)

	// for i := range arr {
	// 	fmt.Println("Array item", i, "is", arr[i])
	// }

}
