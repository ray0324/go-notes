package main

import (
	"fmt"
	"math"
)

type (
	byte int8
)

var (
	m int     = 1000
	n float32 = 1001.01
)

const PI float64 = math.Pi

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

var z complex128 = math.Pi

func main() {
	var a byte = 100
	var b int = math.MaxInt64
	var c = 100
	_, y := 100, "hello"

	var p int = 97
	q := string(p)

	fmt.Println(c / 999)
	fmt.Println(a)
	fmt.Println(math.MaxInt8)
	fmt.Println(b)
	fmt.Println(y)
	fmt.Println(m, int(n))
	fmt.Println(PI)
	fmt.Println(q)
	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)
	fmt.Println(0.1 + 0.2)
	fmt.Println(z)
}
