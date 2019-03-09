package main

import "fmt"

type position struct {
	x, y float64
}

type rect struct {
	width, height float64
	position      position
}

func main() {
	// 1. 基于数组创建切片
	var arr [10]rect = [10]rect{}
	var s1 = arr[0:5]
	// 2. 直接创建数组切片
	s2 := make([]rect, 5, 10)
	// 3.直接创建并初始化切片
	s3 := []rect{{
		width:  100,
		height: 80,
		position: position{
			x: 12,
			y: 24,
		},
	}}

	fmt.Println(s1, s2, s3)
	fmt.Println(s3[0].height)

}
