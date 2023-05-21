package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	// Fibonacii sequence : 0 1 1 2 3 5 8
	go func() {
		var a = 1 // prev element
		var b = 0 // prev of prev element

		ch <- b
		ch <- a

		for i := 2; i < 12; i++ {
			temp := a

			a = a + b
			ch <- (a)
			// fmt.Printf("Writer is sending %d in channel.\n", a+b)
			b = temp
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 12; i++ {
			var f = <-ch
			fmt.Printf("Reader read %d from channel.\n", f)
		}

		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("Process terminated.\n")
}
