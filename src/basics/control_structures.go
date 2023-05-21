package main

import "fmt"

func main() {
	var x int32 = 0

	fmt.Printf("Enter the value of x : ")
	fmt.Scanf("%d", &x)

	if x == 10 {
		fmt.Printf("Value of x is 10.\n")
	} else {
		fmt.Printf("Value of x is not 10.\n")
	}

	switch x {
	case 1:
		fmt.Printf("x is 1.\n")
	case 2:
		fmt.Printf("x is 2.\n")
	default:
		fmt.Printf("x is not 1, 2 or 10")
	}
}
