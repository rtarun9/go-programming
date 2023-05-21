package main

import (
	"fmt"
)

func main() {
	var channel_0_to_1 = make(chan int, 1)
	var channel_1_to_2 = make(chan int, 1)
	var channel_2_to_3 = make(chan int, 1)

	var ending_channel = make(chan string, 1)

	go func() {
		fmt.Printf("Goroutine 1 sending 1\n")
		channel_0_to_1 <- 1
	}()

	go func() {
		message := <-channel_0_to_1
		fmt.Printf("Goroutine 2 recieved :: %d.\n", message)
		fmt.Printf("Goroutine 2 sending 2.\n")
		channel_1_to_2 <- 2
	}()

	go func() {
		message := <-channel_1_to_2
		fmt.Printf("Goroutine 3 recieved :: %d.\n", message)
		fmt.Printf("Goroutine 3 sending 3.\n")
		channel_2_to_3 <- 3
	}()

	go func() {
		message := <-channel_2_to_3
		fmt.Printf("Goroutine 4 recieved :: %d.\n", message)
		fmt.Printf("Goroutine 4 sending 2.\n")
		ending_channel <- "Bye!"
	}()

	end_mesasge := <-ending_channel

	fmt.Printf("ending message :: %s.\n", end_mesasge)
}
