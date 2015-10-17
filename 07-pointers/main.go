package main

import "fmt"

func main() {
	x := 8
	// Pointers are the same as they are in C.
	// The actual memory location is not usually used.
	y := &x
	fmt.Println(x, *y, y)
	*y = 13
	fmt.Println(x, *y, y)

	// new uses function syntax in Go.
	z := new(int)
	*z = 234
	fmt.Println(*z, z)
	// No delete statement is needed, Go uses a garbage collector to clean up memory.
}
