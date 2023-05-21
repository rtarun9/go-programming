package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 10)
	salary_ch := make(chan int, 10)

	go func() {
	updation_loop:
		for {
			designation := <-ch
			if designation == "DONE" {
				fmt.Printf("DONE")
				break updation_loop
			}

			if designation == "AP" {
				salary_ch <- 7500
			}
			if designation == "SAsp" {
				salary_ch <- 12000
			}
		}
	}()

	go func() {
		// Sender
		for i := 0; i < 9; i++ {
			if i%2 == 0 {
				ch <- "AP"
				var salary = <-salary_ch
				fmt.Printf("AP :: %d.\n", salary)
			} else {
				ch <- "SAsp"
				var salary = <-salary_ch
				fmt.Printf("AP :: %d.\n", salary)
			}
		}

		ch <- "DONE"
	}()

	time.Sleep(time.Second * 2)
}
