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
	waitGroup.Add(1)

	go f1(&waitGroup)

	waitGroup.Wait()
	waitGroup.Add(1)

	go f2(&waitGroup)

	waitGroup.Wait()
	waitGroup.Add(1)

	go f3(&waitGroup)

	waitGroup.Wait()
	waitGroup.Add(1)

	go f4(&waitGroup)

	waitGroup.Wait()
	waitGroup.Add(1)

	go f5(&waitGroup)

	waitGroup.Wait()

	fmt.Printf("In the main function!")
}
