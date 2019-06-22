package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

// Args args
type Args struct {
	A, B int
}

// Math math
type Math int

// Multiply multiply
func (m *Math) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Quotient resp struct
type Quotient struct {
	Quo, Rem int
}

// Divide divide
func (m *Math) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("devide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	math := new(Math)
	rpc.Register(math)
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		fmt.Println("Fatal error:", err)
		os.Exit(2)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		fmt.Println("Fatal error:", err)
		os.Exit(2)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("conn error:", err)
			continue
		}

		rpc.ServeConn(conn)

	}

}
