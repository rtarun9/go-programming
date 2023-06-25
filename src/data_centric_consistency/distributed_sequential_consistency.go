package main

import (
	"fmt"
	"sync"
)

type LamportScalarClock struct {
	Time uint32
}

type DataItem struct {
	Value            int32
	LastUpdationTime uint32
}

type Datastore struct {
	X DataItem
	Y DataItem
	Z DataItem

	Mutex sync.Mutex

	Clock LamportScalarClock
}

func (datastore *Datastore) UpdateTime(newTime uint32) {
	datastore.Mutex.Lock()
	defer datastore.Mutex.Unlock()

	if newTime > datastore.Clock.Time {
		datastore.Clock.Time = newTime
	}
}

func (datastore *Datastore) GetCurrentTime() uint32 {
	datastore.Mutex.Lock()
	defer datastore.Mutex.Unlock()

	return datastore.Clock.Time
}

func main() {
	datastore := Datastore{}

	var wg sync.WaitGroup
	wg.Add(3)

	go func(pid int32) {
		defer wg.Done()

		datastore.Mutex.Lock()
		datastore.X.Value = 1
		datastore.X.LastUpdationTime = datastore.GetCurrentTime() + 1
		datastore.UpdateTime(datastore.X.LastUpdationTime)
		datastore.Mutex.Unlock()

		fmt.Printf("[From Processes with ID :: %d] Y = %d, Z = %d.\n", pid, datastore.Y.Value, datastore.Z.Value)
	}(0)

	go func(pid int32) {
		defer wg.Done()

		datastore.Mutex.Lock()
		datastore.Y.Value = 1
		datastore.Y.LastUpdationTime = datastore.GetCurrentTime() + 1
		datastore.UpdateTime(datastore.Y.LastUpdationTime)
		datastore.Mutex.Unlock()

		fmt.Printf("[From Processes with ID :: %d] X = %d, Z = %d.\n", pid, datastore.X.Value, datastore.Z.Value)
	}(1)

	go func(pid int32) {
		defer wg.Done()

		datastore.Mutex.Lock()
		datastore.Z.Value = 1
		datastore.Z.LastUpdationTime = datastore.GetCurrentTime() + 1
		datastore.UpdateTime(datastore.Z.LastUpdationTime)
		datastore.Mutex.Unlock()

		fmt.Printf("[From Processes with ID :: %d] X = %d, Y = %d.\n", pid, datastore.X.Value, datastore.Y.Value)
	}(2)

	wg.Wait()
}
