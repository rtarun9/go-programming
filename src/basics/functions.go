package main

import "fmt"

func area(x int32, y int32) int32 {
	return x * y
}

func main() {
	fmt.Printf("%d", area(10, 20))
}
