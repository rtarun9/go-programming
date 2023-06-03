package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i = i + 1 {
		if i%5 == 0 {
			continue
		}
		fmt.Printf("%d", i)
	}

	for j := 1; j < 10; j = j * 2 {
		fmt.Printf("%f ", (float32)(j))
	}
}
