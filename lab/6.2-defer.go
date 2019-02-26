package main

import "fmt"

func main() {
	// function1()
	f()
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!\n")
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
