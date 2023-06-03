package main

import (
	"fmt"
	"log"
	"net/rpc"
	"strconv"
)

type Reply struct {
	Success bool
	Value   float32
}

func call(client *rpc.Client, _func string, args [2]float32) {
	var reply Reply
	client.Call(_func, args, &reply)
	fmt.Print("Reply :: " + strconv.FormatBool(reply.Success) + " and value :: ")
	fmt.Printf("%f.\n", reply.Value)
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	call(client, "Server.Add", [2]float32{1, 2})
	call(client, "Server.Divide", [2]float32{3, 2})
	call(client, "Server.Add", [2]float32{1, 0})
	call(client, "Server.Divide", [2]float32{1, 0})
}
