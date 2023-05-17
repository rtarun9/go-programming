package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
		return
	}

	var factorialResult int
	err = client.Call("Server.ComputeFactorial", 5, &factorialResult)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result of ComputeFactorial(5):", factorialResult)

	err = client.Call("Server.ComputeFactorial", 3, &factorialResult)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result of ComputeFactorial(3):", factorialResult)

	var powerResult float64
	err = client.Call("Server.ComputeARaisedToB", [2]float64{2.0, 4.0}, &powerResult)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result of ComputeARaisedToB(2.0, 4.0):", powerResult)

	err = client.Call("Server.ComputeARaisedToB", [2]float64{2.0, -1.0}, &powerResult)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result of ComputeARaisedToB(2.0, -1.0):", powerResult)
}
