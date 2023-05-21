package main

import (
	"fmt"
	"time"
)

func main() {
	go func(message string) {
		for i := 0; i < 10; i = i + 1 {
			fmt.Printf("func1 :: %s & %d.\n", message, i)
			time.Sleep(time.Millisecond * 10)
		}
	}("Func1")

	go func(message string) {
		for i := 0; i < 10; i = i + 1 {
			fmt.Printf("func2 :: %s & %d.\n", message, i)
			time.Sleep(time.Millisecond * 10)
		}
	}("Func2")

	time.Sleep(time.Second)
}
