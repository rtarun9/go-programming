package main

import "fmt"

func main() {
	arr := [10]int32{1, 2, 3, 4, 5, 3, 2, 1, 2}
	for i, x := range arr {
		fmt.Printf("%d : %d\n", i, x)
	}

	fmt.Printf("END")

	var arr2 [10]int
	for i, _ := range arr2 {
		arr2[i] = i
	}

	for _, x := range arr2 {
		fmt.Printf("%d ", x)
	}
}
