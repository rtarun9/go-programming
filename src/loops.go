package main

import "fmt"

func main() {
	for i := 0; i < 10; i = i + 1 {
		if i%5 == 0 {
			continue
		}
		fmt.Printf("%d", i)
	}
}
