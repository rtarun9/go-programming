package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wait_group sync.WaitGroup
	wait_group.Add(1)

	go func(message string) {
		for i := 0; i < 10; i = i + 1 {
			fmt.Printf("func1 :: %s & %d.\n", message, i)
			time.Sleep(time.Millisecond * 10)
		}

		wait_group.Done()

	}("Func1")

	go func(message string) {
		wait_group.Wait()

		for i := 0; i < 10; i = i + 1 {
			fmt.Printf("func2 :: %s & %d.\n", message, i)
			time.Sleep(time.Millisecond * 10)
		}
	}("Func2")

	time.Sleep(time.Second * 2)

	fmt.Printf("NEW ------------>")

	var wait_group_2 sync.WaitGroup
	wait_group_2.Add(2)

	go func(message string) {
		for i := 0; i < 10; i = i + 1 {
			fmt.Printf("func1 :: %s & %d.\n", message, i)
			time.Sleep(time.Millisecond * 10)
		}

		wait_group_2.Done()

	}("Func :: X")

	go func(message string) {

		for i := 0; i < 10; i = i + 1 {
			fmt.Printf("func2 :: %s & %d.\n", message, i)
			time.Sleep(time.Millisecond * 10)
		}

		wait_group_2.Done()
	}("Func :: Y ")

	wait_group_2.Wait()
	wait_group_2.Add(2)

	go func(message string) {
		for i := 0; i < 10; i = i + 1 {
			fmt.Printf("func1 :: %s & %d.\n", message, i)
			time.Sleep(time.Millisecond * 10)
		}

		wait_group_2.Done()

	}("Func :: A")

	go func(message string) {

		for i := 0; i < 10; i = i + 1 {
			fmt.Printf("func2 :: %s & %d.\n", message, i)
			time.Sleep(time.Millisecond * 10)
		}

		wait_group_2.Done()
	}("Func :: B")

	wait_group.Wait()
	wait_group_2.Wait()
}
