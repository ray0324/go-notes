package main

import "fmt"

func main() {
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))

	x2, x3 := getX2AndX3_2(2)
	println(x2)
	println(x3)
	// var i1 int = MultiPly3Nums(2, 5, 6)
	// fmt.Printf("MultiPly 2 * 5 * 6 = %d\n", i1)
}

func MultiPly3Nums(a int, b int, c int) int {
	doSth(&a)
	doSth(&a)
	doSth2(a)

	// var product int = a * b * c
	// return product
	return a * b * c
}

func doSth(a *int) {
	b := a
	fmt.Println(b)
	fmt.Println(*b)
}

func doSth2(a int) {
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}
