package main

import (
	"fmt"
	"log"
	"net/rpc"
	"strconv"
)

type Status struct {
	Success bool
	Message string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
		return
	}

	var result Status

	client.Call("Server.InsertEntry", [2]string{"IDK", "I Dont Know"}, &result)
	fmt.Println("Success ::" + strconv.FormatBool(result.Success))
	fmt.Println("Message :: " + result.Message)

	var result2 Status

	client.Call("Server.InsertEntry", [2]string{"JK", "Just Kidding"}, &result2)
	fmt.Println("Success ::" + strconv.FormatBool(result2.Success))
	fmt.Println("Message :: " + result2.Message)

	var result3 Status

	client.Call("Server.InsertEntry", [2]string{"IDK", "I Dont Know"}, &result3)
	fmt.Println("Success ::" + strconv.FormatBool(result3.Success))
	fmt.Println("Message :: " + result3.Message)

	var result4 Status
	client.Call("Server.SearchEntry", "I Dont Know", &result4)
	fmt.Println("Success ::" + strconv.FormatBool(result4.Success))
	fmt.Println("Message :: " + result4.Message)
}
