package main

import "fmt"

// To make a new type, use the type keyword, followed by the name, followed by the
// underlying type.
// The struct type is usually used with a type definition like this.
type foo struct {
	x    int
	y, z string
}

// Any custom type can have functions attached to it like this.
// Just add a receiver between the func keyword and the name.
// You should almost always use a pointer for the receiver.
func (f *foo) printer() {
	fmt.Printf("%d, %s, %s\n", f.x, f.y, f.z)
}

func main() {
	// Structs are made by using their name, followed by their values in braces.
	a := foo{5, "Test", "code"}
	fmt.Println(a)
	x := &a
	// You always use . for accessing things in a struct, even if it is a pointer.
	fmt.Println(x.y)
	x.printer()

	// You can also create a struct using the names of the fields
	b := foo{
		x: 42,
		y: "Foo",
		z: "Bar",
	}
	b.printer()
}
