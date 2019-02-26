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
	// fmt.Println(compare(v1))

	fmt.Printf("Int: %d , Addr: %p\n", v1, ptr)

	// testSwitch(1)
	// testSwitch(2)
	// testSwitch(3)
	// testSwitch(4)

}

// func compare(v1 int) string {
// 	if v1 > 4 {
// 		return "v1>4"
// 	}
// 	return "v1<=4"
// }

func printAddr(p *int) {
	*p = 100
	fmt.Printf(" Addr: %p\n", p)
}

// func testSwitch(i int) {
// 	switch i {
// 	case 1:
// 		fmt.Println("i==1")
// 	case 2, 3:
// 		fmt.Println("i==2,3")
// 	default:
// 		fmt.Println(" not 1,2,3")
// 	}
// }
