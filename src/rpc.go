package main

import (
	"fmt"
	"net"
	"sync"
	"net/rpc"
)

type Server struct {}

func (s *Server) Negate(i int64, reply *int64) error {
	*reply = -i;
	return nil;
}

func server() {
	
	rpc.Register(new(Server));
	ln, err := net.Listen("tcp", ":1235");
	if err != nil {
		fmt.Println(err)
		return;
	}
	for {
		c, err := ln.Accept()
		fmt.Println("Connected to client");
		if err != nil {
			continue
		}
		
		go rpc.ServeConn(c)
	}
}

func client(wg *sync.WaitGroup) {
	
	c, err := rpc.Dial("tcp", "localhost:1235");
	if err != nil {
		fmt.Println(err);
		return;
	}
	
	var result int64;
	err = c.Call("Server.Negate", int64(999), &result);
	if err != nil {
		fmt.Println(err)
		} else {
			fmt.Println("Server.Negate(999) = ", result);
			wg.Done()
		}
	}

func main() {
	var wg sync.WaitGroup;

	wg.Add(1)

	go server()
	go client(&wg)

	wg.Wait();
}