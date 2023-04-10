package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Printf("\nAsync hello world!\n")

	// This will make the program serial essentially.
	// time.Sleep(time.Second)

	fmt.Printf("\nHello World!\n")

	// Removing this will terminate the program since async is now orphan.
	time.Sleep(time.Second)
}
