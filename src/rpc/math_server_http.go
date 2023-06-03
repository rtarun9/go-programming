// Basic math operations performed by server (using HTTP)

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct{}

type Reply struct {
	Success bool
	Value   float32
}

func (server *Server) Add(abPair [2]float32, reply *Reply) error {
	a := abPair[0]
	b := abPair[1]

	reply.Success = true
	reply.Value = a + b

	return nil
}

func (server *Server) Divide(abPair [2]float32, reply *Reply) error {
	a := abPair[0]
	b := abPair[1]

	if b == 0 {
		reply.Success = false
		reply.Value = -1
		return nil
	}

	reply.Success = true
	reply.Value = a / b
	return nil
}

func main() {
	rpc.Register(new(Server))

	rpc.HandleHTTP()

	list, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Print("Server listening @ port 1234")
	http.Serve(list, nil)
}
