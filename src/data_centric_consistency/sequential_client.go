package main

import (
	"fmt"
	"log"
	"net/rpc"
	"sync"
)

type DataItem struct {
	Value            int32
	LastUpdationTime uint32
}

type Datastore struct {
	X DataItem
	Y DataItem
	Z DataItem

	// Mutex is essentially a sync.WaitGroup with internal value as 1.
	Mutex sync.Mutex
}

type Process struct {
	PID uint32

	X DataItem
	Y DataItem
	Z DataItem

	CurrentTime uint32
}

type UpdationRequest struct {
	Variable         byte
	Value            int32
	LastUpdationTime uint32
	Process          Process
}

type ValueRetrievalRequest struct {
	Variable byte
	Process  Process
}

type MultipleValueRetrievalRequest struct {
	Variable1 byte
	Variable2 byte

	Process Process
}

func Process1(wg *sync.WaitGroup) {
	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	p1 := Process{
		PID: 1,
		X: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
		Y: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
		Z: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
	}

	// Update value of x.
	p1.X.Value = 1
	p1.CurrentTime = 1

	updationRequest := UpdationRequest{
		Variable:         'X',
		Value:            1,
		LastUpdationTime: p1.CurrentTime,
		Process:          p1,
	}

	client.Call("Server.UpdateValue", updationRequest, &p1)

	// Get value of Y and Z
	yRetrievalRequest := ValueRetrievalRequest{
		Variable: 'Y',
		Process:  p1,
	}

	zRetrievalRequest := ValueRetrievalRequest{
		Variable: 'Z',
		Process:  p1,
	}

	client.Call("Server.GetValue", yRetrievalRequest, &p1)
	client.Call("Server.GetValue", zRetrievalRequest, &p1)

	fmt.Printf("[From Processes with ID :: %d at local time :: %d] Y = %d, Z = %d.\n", 1, p1.CurrentTime, p1.Y.Value, p1.Z.Value)

	wg.Done()
}

func Process2(wg *sync.WaitGroup) {
	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	p2 := Process{
		PID: 2,
		X: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
		Y: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
		Z: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
	}

	// Update value of y.
	p2.Y.Value = 1
	p2.CurrentTime = 1

	updationRequest := UpdationRequest{
		Variable:         'Y',
		Value:            1,
		LastUpdationTime: p2.CurrentTime,
		Process:          p2,
	}

	client.Call("Server.UpdateValue", updationRequest, &p2)

	// Get value of X and Z
	xRetrievalRequest := ValueRetrievalRequest{
		Variable: 'X',
		Process:  p2,
	}

	zRetrievalRequest := ValueRetrievalRequest{
		Variable: 'Z',
		Process:  p2,
	}

	client.Call("Server.GetValue", xRetrievalRequest, &p2)
	client.Call("Server.GetValue", zRetrievalRequest, &p2)

	fmt.Printf("[From Processes with ID :: %d at local time :: %d] X = %d, Z = %d.\n", 2, p2.CurrentTime, p2.X.Value, p2.Z.Value)

	wg.Done()
}

func Process3(wg *sync.WaitGroup) {
	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	p3 := Process{
		PID: 3,
		X: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
		Y: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
		Z: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
	}

	// Update value of z.
	p3.Z.Value = 1
	p3.CurrentTime = 1

	updationRequest := UpdationRequest{
		Variable:         'Z',
		Value:            1,
		LastUpdationTime: p3.CurrentTime,
		Process:          p3,
	}

	client.Call("Server.UpdateValue", updationRequest, &p3)

	// Get value of X and Y
	xRetrievalRequest := ValueRetrievalRequest{
		Variable: 'X',
		Process:  p3,
	}

	yRetrievalRequest := ValueRetrievalRequest{
		Variable: 'Y',
		Process:  p3,
	}

	client.Call("Server.GetValue", xRetrievalRequest, &p3)
	client.Call("Server.GetValue", yRetrievalRequest, &p3)

	fmt.Printf("[From Processes with ID :: %d at local time :: %d] X = %d, Y = %d.\n", 3, p3.CurrentTime, p3.X.Value, p3.Y.Value)

	wg.Done()
}

// Sequentially consistent version of the above program.

func Process1Consistent(wg *sync.WaitGroup) {
	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	p1 := Process{
		PID: 1,
		X: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
		Y: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
		Z: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
	}

	// Update value of x.
	p1.X.Value = 1
	p1.CurrentTime = 1

	updationRequest := UpdationRequest{
		Variable:         'X',
		Value:            1,
		LastUpdationTime: p1.CurrentTime,
		Process:          p1,
	}

	client.Call("Server.UpdateValue", updationRequest, &p1)
	fmt.Printf("[(Update) From Processes with ID :: %d at local time :: %d] Y = %d, Z = %d.\n", 1, p1.CurrentTime, p1.Y.Value, p1.Z.Value)

	// Get value of Y and Z
	yzRetrievalRequest := MultipleValueRetrievalRequest{
		Variable1: 'Y',
		Variable2: 'Z',
		Process:   p1,
	}

	client.Call("Server.GetValueMultiple", yzRetrievalRequest, &p1)

	fmt.Printf("[(Retrieval) From Processes with ID :: %d at local time :: %d] Y = %d, Z = %d.\n", 1, p1.CurrentTime, p1.Y.Value, p1.Z.Value)

	wg.Done()
}

func Process2Consistent(wg *sync.WaitGroup) {
	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	p2 := Process{
		PID: 2,
		X: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
		Y: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
		Z: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
	}

	// Update value of y.
	p2.Y.Value = 1
	p2.CurrentTime = 1

	updationRequest := UpdationRequest{
		Variable:         'Y',
		Value:            1,
		LastUpdationTime: p2.CurrentTime,
		Process:          p2,
	}

	client.Call("Server.UpdateValue", updationRequest, &p2)
	fmt.Printf("[(Update) From Processes with ID :: %d at local time :: %d] X = %d, Z = %d.\n", 2, p2.CurrentTime, p2.X.Value, p2.Z.Value)

	// Get value of X and Z
	xzRetrievalRequest := MultipleValueRetrievalRequest{
		Variable1: 'X',
		Variable2: 'Z',
		Process:   p2,
	}

	client.Call("Server.GetValueMultiple", xzRetrievalRequest, &p2)

	fmt.Printf("[(Retrieve) From Processes with ID :: %d at local time :: %d] X = %d, Z = %d.\n", 2, p2.CurrentTime, p2.X.Value, p2.Z.Value)

	wg.Done()
}

func Process3Consistent(wg *sync.WaitGroup) {
	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	p3 := Process{
		PID: 3,
		X: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
		Y: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
		Z: DataItem{
			Value:            0,
			LastUpdationTime: 0,
		},
	}

	// Update value of z.
	p3.Z.Value = 1
	p3.CurrentTime = 1

	updationRequest := UpdationRequest{
		Variable:         'Z',
		Value:            1,
		LastUpdationTime: p3.CurrentTime,
		Process:          p3,
	}

	client.Call("Server.UpdateValue", updationRequest, &p3)
	fmt.Printf("[(Update) From Processes with ID :: %d at local time :: %d] X = %d, Y = %d.\n", 3, p3.CurrentTime, p3.X.Value, p3.Y.Value)

	// Get value of X and Y
	xyRetrievalRequest := MultipleValueRetrievalRequest{
		Variable1: 'X',
		Variable2: 'Y',
		Process:   p3,
	}

	client.Call("Server.GetValueMultiple", xyRetrievalRequest, &p3)

	fmt.Printf("[(Retrieve) From Processes with ID :: %d at local time :: %d] X = %d, Y = %d.\n", 3, p3.CurrentTime, p3.X.Value, p3.Y.Value)

	wg.Done()
}

func main() {
	for i := 0; i < 1; i++ {

		wg := sync.WaitGroup{}
		if false {
			fmt.Println("\nSimulation :: Sequentially Inconsistent Data", i)
			wg.Add(3)

			go Process1(&wg)
			go Process2(&wg)
			go Process3(&wg)

			wg.Wait()
		}

		fmt.Println("\nSimulation :: Sequentially Consistent Data", i)

		wg.Add(3)

		go Process1Consistent(&wg)
		go Process2Consistent(&wg)
		go Process3Consistent(&wg)

		wg.Wait()
	}
}
