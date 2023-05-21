package main

import (
	"fmt"
	"time"
)

func f1() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("f1 : %d.\n", i)
	}
}

func f2() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("f2 : %d.\n", i)
	}
}

func main() {
	// This will not be interleaved non-deterministically, since the threads for the go-routines
	// are not launched yet.
	fmt.Printf("\nIn the main function!\n")

	go f1()
	go f2()

	// Inline function.
	go func(message string) {
		fmt.Printf("%s\n", message)
		time.Sleep(100 * time.Millisecond)
	}("Hello")

	time.Sleep(3 * time.Second)
}
