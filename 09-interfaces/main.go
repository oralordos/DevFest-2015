package main

import "fmt"

// Interfaces look like structs with only functions in them.
type printer interface {
	printer() string
}

// New types can be generic types
type foo int

// A type implements an interface if it has all the functions specified in the interface.
func (f foo) printer() string {
	// All of the printing functions have S variants that return the string instead of printing it.
	return fmt.Sprint(f)
}

type bar float32

func (b bar) printer() string {
	return fmt.Sprint(b)
}

// A function can take in an interface like any other type.
// This way I only have to care that the input has the functions I need, nothing else.
func printing(p printer) {
	fmt.Println(p.printer())
}

func main() {
	var a foo = 42
	var b bar = 3.14
	printing(a)
	printing(b)
}
