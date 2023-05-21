package main

import (
	"fmt"
	"log"
	"net/rpc"
	"strconv"
	"time"
)

type RoomType struct {
	TypeIndex       int16 // 0 indicate a single room, 1 indicates double room, etc.
	CostPerNight    int64
	TotalRoomCount  int32
	BookedRoomCount int32
}

type Status struct {
	Message string
	Success bool
	Time    time.Time
}

type AvailabilityQueryResponse struct {
	Status_   Status
	RoomData_ RoomType
}

func PrintRoomType(roomType *RoomType) {
	fmt.Printf("RoomType:\n")
	fmt.Printf("\tTypeIndex: %d\n", roomType.TypeIndex)
	fmt.Printf("\tCostPerNight: Rs.%d\n", roomType.CostPerNight)
	fmt.Printf("\tTotalRoomCount: %d\n", roomType.TotalRoomCount)
	fmt.Printf("\tBookedRoomCount: %d\n\n", roomType.BookedRoomCount)
}

func PrintStatus(status *Status) {
	fmt.Print("Status :: \n\tMessage -> " + status.Message + ".\n\tSuccess -> " + strconv.FormatBool(status.Success) + ".\n\tTime -> " + status.Time.String() + ".\n\n")
}

func PrintAvailabilityQueryResponse(response *AvailabilityQueryResponse) {

	PrintStatus(&response.Status_)
	fmt.Printf("\n")
	if response.Status_.Success == true {
		PrintRoomType(&response.RoomData_)
	}
	fmt.Printf("\n")
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error connecting to the server:", err)
	}

	// Check for availability of rooms.
	var availabilityQueryReponse AvailabilityQueryResponse
	err = client.Call("Server.CheckRoomAvailability", 0, &availabilityQueryReponse)
	if err != nil {
		log.Fatal("Error while checking for availability of room with index 0.", err)
	}

	fmt.Printf("Server.PrintAvailabilityQueryResponse(0) = \n")
	PrintAvailabilityQueryResponse(&availabilityQueryReponse)

	// ------------------------------------------------------------------------------

	err = client.Call("Server.CheckRoomAvailability", 5, &availabilityQueryReponse)
	if err != nil {
		log.Fatal("Error while checking for availability of room with index 5.", err)
	}

	fmt.Printf("Server.PrintAvailabilityQueryResponse(5) = \n")
	PrintAvailabilityQueryResponse(&availabilityQueryReponse)

	// ------------------------------------------------------------------------------

	// Book a few rooms.
	var bookingRequestResponse Status
	err = client.Call("Server.BookRoom", 0, &bookingRequestResponse)
	if err != nil {
		log.Fatal("Error while booking a room:", err)
	}

	fmt.Printf("Server.BookRoom(0) = \n")
	PrintStatus(&bookingRequestResponse)

	// ------------------------------------------------------------------------------

	err = client.Call("Server.BookRoom", 3, &bookingRequestResponse)
	if err != nil {
		log.Fatal("Error while booking a room:", err)
	}

	fmt.Printf("Server.BookRoom(3) = \n")
	PrintStatus(&bookingRequestResponse)

	// ------------------------------------------------------------------------------

	err = client.Call("Server.BookRoom", 9, &bookingRequestResponse)
	if err != nil {
		log.Fatal("Error while booking a room:", err)
	}

	err = client.Call("Server.CheckRoomAvailability", 3, &availabilityQueryReponse)
	if err != nil {
		log.Fatal("Error while checking for availability of room with index 3.", err)
	}

	fmt.Printf("Server.PrintAvailabilityQueryResponse(3) = \n")
	PrintAvailabilityQueryResponse(&availabilityQueryReponse)

	// ------------------------------------------------------------------------------

	// Perform some room cancellations.
	var cancellationRequestReponse Status
	err = client.Call("Server.CancelRoomBooking", 0, &cancellationRequestReponse)
	if err != nil {
		log.Fatal("Error while cancelling room booking of room with index 0.", err)
	}

	fmt.Printf("Server.CancelRoomBooking(0) = \n")
	PrintStatus(&cancellationRequestReponse)

	// ------------------------------------------------------------------------------

	err = client.Call("Server.CancelRoomBooking", 3, &cancellationRequestReponse)
	if err != nil {
		log.Fatal("Error while cancelling room booking of room with index 3.", err)
	}

	fmt.Printf("Server.CancelRoomBooking(3) = \n")
	PrintStatus(&cancellationRequestReponse)

	// ------------------------------------------------------------------------------

	err = client.Call("Server.CancelRoomBooking", 2, &cancellationRequestReponse)
	if err != nil {
		log.Fatal("Error while cancelling room booking of room with index 2.", err)
	}

	fmt.Printf("Server.CancelRoomBooking(2) = \n")
	PrintStatus(&cancellationRequestReponse)
}
