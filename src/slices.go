package main

import "fmt"

func display(x []int) {
	for index, elem := range x {
		fmt.Printf("index %d :: elem %d\n", index, elem)
	}
}

func main() {
	// Slices are like arrays, but are dynamic.

	// a array.
	var a [5]int = [5]int{5, 6, 7, 8, 9}
	fmt.Println(a)

	// a slice.
	var b []int = []int{3, 2, 1}
	b = append(b, 111, 222, 333, 444, 555)
	display(b)
}
