package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 2)

	go func() {
		fmt.Printf("\t[INTERNAL] :: Function 1 Sending 10\n")
		ch1 <- 10
	}()

	go func() {
		fmt.Printf("\t[INTERANL] :: Function 2 sending 20\n")
		ch2 <- 20
	}()

	time.Sleep(time.Millisecond * 2)

	select {
	case c1 := <-ch1:
		fmt.Printf("channel 1 send the value :: %d.\n", c1)

	case c2 := <-ch2:
		fmt.Printf("channel 2 send the value :: %d.\n", c2)
	default:
		fmt.Printf("No channel sent message :(")
	}
}
