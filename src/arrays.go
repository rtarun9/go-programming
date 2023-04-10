package main

import "fmt"

func main() {
	arr := [10]int32{1, 2, 3, 4, 5, 3, 2, 1, 2}
	for i, x := range arr {
		fmt.Printf("%d : %d ", i, x)
	}
}
