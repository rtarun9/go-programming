package main

import (
	"fmt"
	"log"
	"net/rpc"
	"sync"
)

func client(i int, wg *sync.WaitGroup) {
		client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
		return
	}

	var factorialResult int
	err = client.Call("Server.ComputeFactorial", i, &factorialResult)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("CLIENT :: ", i, " -> Result of ComputeFactorial(", i, ") :: ", factorialResult)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(10)

	for x := 0; x < 10; x++ {
		go client(x, &wg)
	}

	wg.Wait()
}
