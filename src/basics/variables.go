package main

import "fmt"

func main() {
	var a int = 20
	var b string
	b = "Hello, World!" // = is a assignment.
	var c = 133.30

	x := 333 // declaration and assignment with inference.

	t := 3 // There is no need to put = here, a := must be used.

	fmt.Printf("a is %d.\n", a)
	fmt.Printf("b is %s.\n", b)
	fmt.Printf("c is %f.\n", c)
	fmt.Printf("x is %d.\n", x)
	fmt.Printf("t is %d.\n", t)

	var y int
	x, y = get_data(false)
	dx, dy := get_data(true)
	fmt.Printf("x is %d.\ny is %d.\n", x, y)
	fmt.Printf("dx is %d.\ndy is %d.\n", dx, dy)

	var _a uint8 = uint8('a')
	fmt.Println(_a)
}

func get_data(branch bool) (int, int) {
	if branch == true {
		return -32, 32
	} else {
		return 111, -111
	}
}
