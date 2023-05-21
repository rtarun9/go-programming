package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 9)
	var wg sync.WaitGroup

	go func() {
		// Sender process.
		defer close(ch)

		a := 1
		b := 0

		ch <- b
		ch <- a

	sender_loop:
		for {
			temp := a
			a = a + b
			select {
			case ch <- a:
				b = temp
			default:
				ch <- -1
				break sender_loop
			}
		}
	}()

	wg.Add(1)

	go func() {
		// Reader process

	receiver_loop:
		for {
			var fib = 0

			select {
			case fib = <-ch:
				if fib == -1 {
					wg.Done()
					break receiver_loop
				}
				fmt.Printf("%d ", fib)
			}
		}
	}()

	wg.Wait()
}
