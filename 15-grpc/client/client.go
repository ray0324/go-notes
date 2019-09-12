package main

import (
	"fmt"
	"log"
	"net/rpc"

	"../defs"
)

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Println("Usage:", os.Args[0], "server")
	// 	os.Exit(1)
	// }

	// serverAddr := os.Args[1]

	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := defs.Args{A: 17, B: 8}

	var reply int

	err = client.Call("Math.Multiply", args, &reply)

	if err != nil {
		log.Fatal("Math error:", err)
	}

	fmt.Printf("Math %d * %d = %d\n", args.A, args.B, reply)

	var quo defs.Quotient

	err = client.Call("Math.Divide", args, &quo)

	if err != nil {
		log.Fatal("Math error:", err)
	}

	fmt.Printf("Math %d / %d = %d ..... %d\n", args.A, args.B, quo.Quo, quo.Rem)

}
