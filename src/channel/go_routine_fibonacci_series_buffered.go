package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(1)

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
				wg.Done()
				fmt.Print("SENDER IS DONE")
				break sender_loop
			}
		}
	}()

	time.Sleep(time.Second * 1)
	wg.Wait()
	wg.Add(1)
	go func() {
		// Reader process

	receiver_loop:
		for {
			var fib = 0

			select {
			case fib = <-ch:
				fmt.Printf("%d ", fib)
				fmt.Println(len(ch))
			default:
				wg.Done()
				break receiver_loop
			}
		}
	}()

	wg.Wait()
}
