package main

import "fmt"

var array = [3]int{100, 200, 300}
var arr2 = [2][3]int{{100, 200, 300}, {100, 200, 300}}

// Tmp1 test pointer
func Tmp1() {

	// 数组迭代
	for _, value := range arr2 {
		for i, v := range value {
			fmt.Println(i, v)
		}
	}

	fmt.Println("=========")
	// 数组作为函数参数  传递的是数组的副本
	optArray(array)

	fmt.Println(array)
	fmt.Println(&array[0])
	fmt.Println("=========")

	// 操作函数指针
	optArrayPtr(&array)

	fmt.Println(array)
	fmt.Println(&array[0])
	fmt.Println("=========")

	fmt.Println(arr2)
}

func optArray(arr [3]int) {
	fmt.Println(&arr)
	fmt.Println(&arr[0])
	arr[2] = 600
}

func optArrayPtr(arr *[3]int) {

	fmt.Println(&arr)
	fmt.Println(arr)
	fmt.Println(&arr[0])
	fmt.Println(arr[0])
	arr[2] = 600

}
