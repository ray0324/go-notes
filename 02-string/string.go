package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var str = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")

	// 字符串是否以指定的prefix开头
	fmt.Printf("\n%t\n", strings.HasPrefix(str, "Th"))

	// 查找字符串的索引
	fmt.Printf("%d\n", strings.Index(str, "string"))

	fmt.Printf("%d\n", strings.LastIndex(str, "an"))
	// 字符串替换
	fmt.Printf("\n%s\n", strings.Replace(str, "is", "at", -1))
	// 统计字符串出现的次数
	fmt.Printf("\n%d\n", strings.Count(str, "is"))

	// 重复特定的字符串
	fmt.Printf("%s\n", strings.Repeat("A", 10))
	// 转换大小写
	fmt.Printf("%s\n", strings.ToLower(str))
	fmt.Printf("%s\n", strings.ToUpper(str))

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// integer to ascii
	fmt.Printf("The size of ints is: %s\n", strconv.Itoa(1000))
	fmt.Printf("The size of ints is: %s\n", strconv.FormatFloat(math.Pi, 'e', 16, 64))
	// ascii to integer
	num, err := strconv.Atoi("12345a")
	fmt.Printf("The size of ints is: %d\n %s\n", num, err)

}
