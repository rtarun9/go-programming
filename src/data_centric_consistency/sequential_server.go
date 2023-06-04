package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

type Server struct{}

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

var datastore = Datastore{}

func (server *Server) UpdateValue(request UpdationRequest, response *Process) error {
	log.Println("Updation Request from :: ", request.Process.PID)

	for !datastore.Mutex.TryLock() {
		request.Process.CurrentTime += 1
	}

	if request.Variable == 'X' {
		datastore.X.Value = request.Value
		datastore.X.LastUpdationTime = request.Process.CurrentTime + 1
	} else if request.Variable == 'Y' {
		datastore.Y.Value = request.Value
		datastore.Y.LastUpdationTime = request.Process.CurrentTime + 1
	} else if request.Variable == 'Z' {
		datastore.Z.Value = request.Value
		datastore.Z.LastUpdationTime = request.Process.CurrentTime + 1
	}

	*response = request.Process

	fmt.Printf("State of global datastore: \n\tX value and update time : %d %d\n\tY value and update time : %d %d\n\tZ Value and update time : %d %d\n",
		datastore.X.Value, datastore.X.LastUpdationTime, datastore.Y.Value, datastore.Y.LastUpdationTime, datastore.Z.Value, datastore.Z.LastUpdationTime)

	datastore.Mutex.Unlock()

	return nil
}

func (server *Server) GetValue(request ValueRetrievalRequest, response *Process) error {
	log.Println("GetValue Request from :: ", request.Process.PID)

	for !datastore.Mutex.TryLock() {
		request.Process.CurrentTime += 1
	}

	if request.Variable == 'X' {
		if request.Process.CurrentTime < datastore.X.LastUpdationTime {
			request.Process.X = datastore.X
		}
	} else if request.Variable == 'Y' {
		if request.Process.CurrentTime < datastore.Y.LastUpdationTime {
			request.Process.Y = datastore.Y
		}
	} else if request.Variable == 'Z' {
		if request.Process.CurrentTime < datastore.Z.LastUpdationTime {
			request.Process.Z = datastore.Z
		}
	}

	request.Process.CurrentTime += 1

	*response = request.Process

	fmt.Printf("State of global datastore: \n\tX value and update time : %d %d\n\tY value and update time : %d %d\n\tZ Value and update time : %d %d\n",
		datastore.X.Value, datastore.X.LastUpdationTime, datastore.Y.Value, datastore.Y.LastUpdationTime, datastore.Z.Value, datastore.Z.LastUpdationTime)
	datastore.Mutex.Unlock()

	return nil
}

func main() {
	rpc.Register(new(Server))

	listener, error := net.Listen("tcp", "localhost:1234")
	if error != nil {
		log.Fatal(error)
		log.Fatal("ERROR~")
		return
	}

	log.Println("Server listening at localhost:1234 (TCP)")
	for {
		conn, error := listener.Accept()
		if error != nil {
			log.Fatal(error)
			return
		}

		fmt.Print("Server connected to client!\n\n")
		go rpc.ServeConn(conn)
	}
}
