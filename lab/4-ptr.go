package main

import "fmt"

func main() {
	var ptr *int
	var p **int
	var q ***int
	num := 1000
	ptr = &num

	p = &ptr
	q = &p

	fmt.Println(&ptr, p, &num, ptr, **p, ***q, *ptr)

}
