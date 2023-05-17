// The server has a few services it provides:
// (i) Compute the factorial of a number.
// (ii) Compute the power of a raised to b.

package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Server struct{}

func (server *Server) ComputeFactorial(n int, reply *int) error {
	var result = 1
	for i := n; i >= 1; i-- {
		result *= i
	}

	*reply = result
	return nil
}

func (server *Server) ComputeARaisedToB(abPair [2]float64, reply *float64) error {
	var a = abPair[0]
	var b = abPair[1]

	var result = 1.0

	if b < 0.0 {
		b = -1.0 * b
		a = 1.0 / a
	}
	for i := 0; i < int(b); i++ {
		result *= a
	}
	*reply = result
	return nil
}

func main() {
	rpc.Register(new(Server))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Server is listening on port 1234")

	for {
		conn, err := listener.Accept()
		fmt.Println("Connected to client!")
		if err != nil {
			continue
		}

		go rpc.ServeConn(conn)
	}
}
