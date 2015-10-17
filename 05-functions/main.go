package main

import "fmt"

// Constants can be defined using const instead of var.
const value = 7

// Functions start with keywork func, then the name, then the parameter list, then the return type.
// The type is after the parameter name.
func plusSix(x int) int {
	return x + 6
}

func main() {
	printer("Foo", "Bar")
	fmt.Println(value, "+ 6 =", plusSix(value))
	x := "Closures work in Go!"
	f := func() {
		fmt.Println(x)
	}
	f()
	more()
	d, s := doubleAndSize("This is text")
	fmt.Printf("Double is %s, size is %d\n", d, s)
}

// Multiple parameters of the same type can collapse their types together.
// Functions can be placed anywhere in the file.
// If there is no return type, just skip putting anything there.
func printer(val1, val2 string) {
	fmt.Println(val1, val2)
}

// Multiple return types are also allowed.
func doubleAndSize(val string) (string, int) {
	x := val + val
	return x, len(x)
}
