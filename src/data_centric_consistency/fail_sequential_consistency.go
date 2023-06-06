// The idea is for each process to have a local time.
// The data store is global and has a variable storing the last time
// the variable was updated.
package main

import (
	"fmt"
	"math"
	"sync"
)

type DataItem struct {
	Value            int32
	LastUpdationTime uint32
}

type Datastore struct {
	X DataItem
	Y DataItem
	Z DataItem

	// Mutex is essentially a sync.WaitGroup with internal value as 1.
	Mutex sync.Mutex
}

type Process struct {
	PID uint32

	X DataItem
	Y DataItem
	Z DataItem

	CurrentTime uint32
}

func main() {
	// Running the simulation for 4 times.
	for i := 0; i < 4; i++ {
		fmt.Println("\nSimulation :: ", i)

		datastore := Datastore{}

		p1 := Process{
			PID: 1,
			X: DataItem{
				Value:            0,
				LastUpdationTime: 0,
			},
			Y: DataItem{
				Value:            0,
				LastUpdationTime: 0,
			},
			Z: DataItem{
				Value:            0,
				LastUpdationTime: 0,
			},
		}

		p2 := Process{
			PID: 2,
			X: DataItem{
				Value:            0,
				LastUpdationTime: 0,
			},
			Y: DataItem{
				Value:            0,
				LastUpdationTime: 0,
			},
			Z: DataItem{
				Value:            0,
				LastUpdationTime: 0,
			},
		}

		p3 := Process{
			PID: 3,
			X: DataItem{
				Value:            0,
				LastUpdationTime: 0,
			},
			Y: DataItem{
				Value:            0,
				LastUpdationTime: 0,
			},
			Z: DataItem{
				Value:            0,
				LastUpdationTime: 0,
			},
		}

		wg := sync.WaitGroup{}
		wg.Add(3)

		go func(process Process) {
			process.X.Value = 1
			process.CurrentTime = 1

			for !datastore.Mutex.TryLock() {
				process.CurrentTime += 1
			}

			// Get the value of Y and Z.
			if process.CurrentTime < datastore.Z.LastUpdationTime {
				process.Z = datastore.Z
			}

			process.CurrentTime = uint32(math.Max(float64(process.CurrentTime)+1, math.Max(float64(datastore.Y.LastUpdationTime)+1, float64(datastore.Z.LastUpdationTime)+1)))
			process.X.LastUpdationTime = process.CurrentTime

			datastore.X = process.X
			datastore.Mutex.Unlock()

			// Update last updated value of X in datastore.
			fmt.Printf("[From Processes with ID :: %d at local time :: %d] Y = %d, Z = %d.\n", process.PID, process.CurrentTime, process.Y.Value, process.Z.Value)

			wg.Done()
		}(p1)

		go func(process Process) {
			process.Y.Value = 1
			process.CurrentTime = 1

			// Get the value of X and Z.
			for !datastore.Mutex.TryLock() {
				process.CurrentTime += 1
			}

			if process.CurrentTime < datastore.Z.LastUpdationTime {
				process.Z = datastore.Z
			}

			process.CurrentTime = uint32(math.Max(float64(process.CurrentTime)+1, math.Max(float64(datastore.X.LastUpdationTime)+1, float64(datastore.Z.LastUpdationTime)+1)))
			process.Y.LastUpdationTime = process.CurrentTime

			datastore.Y = process.Y
			datastore.Mutex.Unlock()

			// Update last updated value of X in datastore.
			fmt.Printf("[From Processes with ID :: %d at local time :: %d] X = %d, Z = %d.\n", process.PID, process.CurrentTime, process.X.Value, process.Z.Value)

			wg.Done()
		}(p2)

		go func(process Process) {
			process.Z.Value = 1
			process.CurrentTime = 1

			// Get the value of X and Y.
			for !datastore.Mutex.TryLock() {
				process.CurrentTime += 1
			}

			if process.CurrentTime < datastore.Y.LastUpdationTime {
				process.Y = datastore.Y
			}

			process.CurrentTime = uint32(math.Max(float64(process.CurrentTime)+1, math.Max(float64(datastore.X.LastUpdationTime)+1, float64(datastore.Y.LastUpdationTime)+1)))
			process.Z.LastUpdationTime = process.CurrentTime

			datastore.Z = process.Z
			datastore.Mutex.Unlock()

			// Update last updated value of X in datastore.
			fmt.Printf("[From Processes with ID :: %d at local time :: %d] X = %d, Y = %d.\n", process.PID, process.CurrentTime, process.X.Value, process.Y.Value)

			wg.Done()
		}(p3)

		wg.Wait()

		fmt.Printf("State of global datastore: \n\tX value and update time : %d %d\n\tY value and update time : %d %d\n\tZ Value and update time : %d %d\n",
			datastore.X.Value, datastore.X.LastUpdationTime, datastore.Y.Value, datastore.Y.LastUpdationTime, datastore.Z.Value, datastore.Z.LastUpdationTime)
	}
}
