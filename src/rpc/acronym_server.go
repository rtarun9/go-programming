package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strconv"
)

type Server struct{}

type Status struct {
	Success bool
	Message string
}

var acronym_db = make(map[string]string)

func (server *Server) InsertEntry(request [2]string, result *Status) error {
	acronym := request[0]
	word := request[1]

	_, ok := acronym_db[word]
	fmt.Print(strconv.FormatBool(ok))
	if ok {
		result.Success = false
		result.Message = "Key already exist in acronym database!"
		return nil
	}

	result.Success = true
	result.Message = "Succesfully entered " + string(acronym) + " :: " + string(word) + " into the database.\n"

	acronym_db[word] = acronym

	return nil
}

func (server *Server) SearchEntry(word string, result *Status) error {
	value, ok := acronym_db[word]
	if ok {
		result.Success = true
		result.Message = value
		return nil
	}

	result.Message = "ERROR"
	result.Success = false
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

	fmt.Println("Server listening at localhost:1234 (TCP)")
	for {
		conn, error := listener.Accept()
		if error != nil {
			log.Fatal(error)
			return
		}

		fmt.Print("Server connected to client!")
		go rpc.ServeConn(conn)
	}
}
