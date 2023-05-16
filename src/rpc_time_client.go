package main

import (
	"net/rpc"
	"log"
)

type Args struct {}

type TimeServer int64

func main() {
	var reply int64
	args := Args{}

	server, err := rpc.DialHTTP("tcp", "0.0.0.0:1234");
	if err != nil {;
		log.Fatal("Client to server connection error.")
	}

	err = server.Call("TimeServer.GiveServerTime", args, &reply)
	if err != nil {
		log.Fatal("ERROR")
	} else {
		log.Printf("Reply :: %d", reply);
	}
}