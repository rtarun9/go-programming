package main

import "fmt"

func main() {

	// Create a channel. The expected data type that will travel through the channel is 'string'.
	messages := make(chan string)

	// Inline function.
	go func(msg string) {
		// Send the value of "ping" through the channel, reciever is blocked until the value is recieved.
		messages <- "ping"
		fmt.Println(msg)
	}("In func")

	// Explicit lock, untill the data is recieved by the reciever.
	msg := <-messages
	fmt.Println(msg)
}
