package main

import "fmt"

func main() {
	// This is how you make an array in Go.
	data := [3]string{"Test", "Golang", "Google"}

	// This prints the type of a variable.
	fmt.Printf("data is of the type %T\n", data)

	// The standard for loop from C is available.
	for i := 0; i < len(data); i++ {
		fmt.Printf("%d-%s ", i, data[i])
	}
	fmt.Println()

	// Looping over an array is extremely common, so range for loops appear a lot.
	for i := range data {
		fmt.Printf("%d-%s ", i, data[i])
	}
	fmt.Println()

	// It is rare that you need just the index in a range loop, so the value is available as well.
	for i, v := range data {
		fmt.Printf("%d-%s ", i, v)
	}
	fmt.Println()

	// Go uses for loops even for while statements, just put only a condition and it works.
	i := 0
	for i < len(data) {
		fmt.Printf("%d-%s ", i, data[i])
		i++
	}
	fmt.Println()

	// If you skip everything, it will make an infinite loop.
	// for {
	//   fmt.Println("This is an infinite loop.")
	// }
}
