// The problem is basically a pipeline where each stage depends on the output
// of the previous stage. The goal is to have several go-routines
// run in a serial manner.
package main

import (
	"fmt"
	"sync"
	"time"
)

func EngineSetup(waitGroup *sync.WaitGroup) {
	fmt.Printf("..Engine setup...\n")
	time.Sleep(time.Second * 1)
	waitGroup.Done()
}

func AcceleratorBrakeClutchSetup(waitGroup *sync.WaitGroup) {
	fmt.Printf("..Accelerator-Brake-Clutch setup...\n")
	time.Sleep(time.Second * 1)
	waitGroup.Done()
}

func BodyAssemblingStage(waitGroup *sync.WaitGroup) {
	fmt.Printf("..Body-Assembling setup...\n")
	time.Sleep(time.Second * 1)
	waitGroup.Done()
}

func SteeringSetup(waitGroup *sync.WaitGroup) {
	fmt.Printf("..Steering setup...\n")
	time.Sleep(time.Second * 1)
	waitGroup.Done()
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go EngineSetup(&waitGroup)
	waitGroup.Wait()
	waitGroup.Add(1)

	go AcceleratorBrakeClutchSetup(&waitGroup)
	waitGroup.Wait()
	waitGroup.Add(1)

	go BodyAssemblingStage(&waitGroup)
	waitGroup.Wait()
	waitGroup.Add(1)

	go SteeringSetup(&waitGroup)

	waitGroup.Wait()
	fmt.Print("Car is now completely setup!\n")

	return
}
