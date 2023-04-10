package main

import (
	"fmt"
	"sync"
)

func f1(waitGroup *sync.WaitGroup) {
	fmt.Printf("F1!\n")
	waitGroup.Done()
}

func f2(waitGroup *sync.WaitGroup) {
	fmt.Printf("F2!\n")
	waitGroup.Done()
}

func f3(waitGroup *sync.WaitGroup) {
	fmt.Printf("F3!\n")
	waitGroup.Done()
}

func f4(waitGroup *sync.WaitGroup) {
	fmt.Printf("F4!\n")
	waitGroup.Done()
}

func f5(waitGroup *sync.WaitGroup) {
	fmt.Printf("F5!\n")
	waitGroup.Done()
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(5)

	go f1(&waitGroup)
	go f2(&waitGroup)
	fmt.Printf("\nIn main, between several go routine invocations\n")
	go f3(&waitGroup)
	go f4(&waitGroup)
	go f5(&waitGroup)

	waitGroup.Wait()

	fmt.Printf("In the main function!")
}
