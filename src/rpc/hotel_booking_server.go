package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type Server struct{}

type RoomType struct {
	TypeIndex       int16 // 0 indicate a single room, 1 indicates double room, etc.
	CostPerNight    int64
	TotalRoomCount  int32
	BookedRoomCount int32
}

func GetRoomTypeName(roomTypeIndex int16) string {
	if roomTypeIndex == 0 {
		return "Single Room"
	} else if roomTypeIndex == 1 {
		return "Double Room"
	} else if roomTypeIndex == 2 {
		return "Twin Room"
	} else if roomTypeIndex == 3 {
		return "Triple Room"
	} else if roomTypeIndex == 4 {
		return "Quad Room"
	}

	log.Fatal("INVALID ROOT TYPE INDEX")
	return "INVALID ROOM TYPE"
}

var rooms [5]RoomType

type Status struct {
	Message string
	Success bool
	Time    time.Time
}

type AvailabilityQueryResponse struct {
	Status_   Status
	RoomData_ RoomType
}

func FillStatus(status *Status, message string, success bool) {
	status.Message = message
	status.Success = success
	status.Time = time.Now()
}

// Funcionality 1 :: Client can book a room of type available.
func (server *Server) BookRoom(roomTypeIndex int, bookingRequestResponse *Status) error {
	// Check of roomTypeIndex is in range.
	if roomTypeIndex < 0 || roomTypeIndex > 4 {
		FillStatus(bookingRequestResponse, "Invalid Room Type Index. Must be in range 0..4", false)
		return nil
	}

	// Check if any rooms are availalble.
	if rooms[roomTypeIndex].BookedRoomCount >= rooms[roomTypeIndex].TotalRoomCount {
		FillStatus(bookingRequestResponse, "No available rooms. Sorry for the inconvinience.", false)
		return nil
	}

	// If code reaches this point, a room can be booked.
	FillStatus(bookingRequestResponse, "Succesfully booking your room! Type : "+GetRoomTypeName(int16(roomTypeIndex)), true)

	rooms[roomTypeIndex].BookedRoomCount++
	return nil
}

// Funcionality 2 :: Client can cancel booking of a room of type.
func (server *Server) CancelRoomBooking(roomTypeIndex int, cancellationRequestResponse *Status) error {
	// Check of roomTypeIndex is in range.
	if roomTypeIndex < 0 || roomTypeIndex > 4 || roomTypeIndex > len(rooms) {
		FillStatus(cancellationRequestResponse, "Invalid Room Type Index. Must be in range 0..4", false)
		return nil
	}

	// Check if any rooms were booked in the first place.
	if rooms[roomTypeIndex].BookedRoomCount <= 0 {
		FillStatus(cancellationRequestResponse, "None of the rooms of this type were booked. Was there a misinput on your side? Please check and try again. Room type : "+GetRoomTypeName(int16(roomTypeIndex)), false)
		return nil
	}

	// If code reaches this point, the room can be cancelled.
	FillStatus(cancellationRequestResponse, "Succesfully cancelled your room of type: "+GetRoomTypeName(int16(roomTypeIndex)), true)

	rooms[roomTypeIndex].BookedRoomCount--
	return nil
}

// Funcionality 3 :: Client can check status of room (availability).
func (server *Server) CheckRoomAvailability(roomTypeIndex int, availabilityRequestResponse *AvailabilityQueryResponse) error {
	// Check of roomTypeIndex is in range.
	if roomTypeIndex < 0 || roomTypeIndex > 4 || roomTypeIndex > len(rooms) {
		FillStatus(&availabilityRequestResponse.Status_, "Invalid Room Type Index. Must be in range 0..4", false)
		availabilityRequestResponse.RoomData_.BookedRoomCount = -1
		availabilityRequestResponse.RoomData_.CostPerNight = -1
		availabilityRequestResponse.RoomData_.TotalRoomCount = -1
		availabilityRequestResponse.RoomData_.TypeIndex = int16(roomTypeIndex)
		return nil
	}

	availabilityRequestResponse.RoomData_ = rooms[roomTypeIndex]
	FillStatus(&availabilityRequestResponse.Status_, "Succesfully queried availability of room of type: "+GetRoomTypeName(int16(roomTypeIndex)), true)
	return nil
}

func main() {
	// Setup initial data of rooms as given in question.
	rooms[0].TypeIndex = 0
	rooms[0].TotalRoomCount = 10
	rooms[0].BookedRoomCount = 0
	rooms[0].CostPerNight = 1000

	rooms[1].TypeIndex = 1
	rooms[1].TotalRoomCount = 20
	rooms[1].BookedRoomCount = 0
	rooms[1].CostPerNight = 1500

	rooms[2].TypeIndex = 2
	rooms[2].TotalRoomCount = 5
	rooms[2].BookedRoomCount = 0
	rooms[2].CostPerNight = 2000

	rooms[3].TypeIndex = 3
	rooms[3].TotalRoomCount = 3
	rooms[3].BookedRoomCount = 0
	rooms[3].CostPerNight = 3000

	rooms[4].TypeIndex = 4
	rooms[4].TotalRoomCount = 2
	rooms[4].BookedRoomCount = 0
	rooms[4].CostPerNight = 5000

	rpc.Register(new(Server))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Server is listening on port 1234")

	for {
		conn, err := listener.Accept()
		fmt.Println("Connected to client!")
		if err != nil {
			continue
		}

		go rpc.ServeConn(conn)
	}
}
