package main

import "fmt"

var ptr *int

func main() {

	var v1 = 5

	ptr = &v1

	fmt.Println(v1)

	*ptr = 3

	printAddr(ptr)
	fmt.Println(v1)
	fmt.Printf("Int: %d , Addr: %p\n", v1, ptr)
	testptr()

}

func printAddr(p *int) {
	*p = 100
	fmt.Printf(" Addr: %p\n", p)
}

func testptr() {
	var ptr *int
	var p **int
	var q ***int
	num := 1000
	ptr = &num

	p = &ptr
	q = &p

	fmt.Println(&ptr, p, &num, ptr, **p, ***q, *ptr)
}
