package main

import (
	"time"
	"net"
	"net/rpc"
	"net/http"
	"log"
)

type Args struct {}

type TimeServer int64

func main() {
	timeserver := new(TimeServer)

	rpc.Register(timeserver)

	rpc.HandleHTTP()
	list, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listner error")
	}

	http.Serve(list, nil)
}

func (t* TimeServer) GiveServerTime(args *Args, reply *int64) error {
	*reply = time.Now().Unix()
	return nil

}