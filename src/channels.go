package main

import "fmt"

func main() {
	message := make(chan int, 3)

	go func() {
		// Send to channel.
		message <- 1
	}()

	go func() {
		// Send to channel.
		message <- 2
	}()

	go func() {
		// Send to channel.
		message <- 3
	}()

	var msg int = <-message
	fmt.Printf("Message received from go-routine :: %d.\n", msg)

	msg = <-message
	fmt.Printf("Message received from go-routine :: %d.\n", msg)

	msg = <-message
	fmt.Printf("Message received from go-routine :: %d.\n", msg)

}
