// First line must be a package statement.
// package will be main for an executable, something else if it is meant to be imported.
package main

// Import statements work in a similar way to python's imports.
import "fmt"

// package main function main is the entry point into the program.
func main() {
	// The Println function is from the fmt package, it is the normal way to output text in Go.
	fmt.Println("Hello World!")
	// The Print function does the same thing as Println, but without the newline at the end.
	fmt.Print("What is your name? ")
	// This is how you define an empty variable, var <name> <type>
	var myName string
	// This is how you get input, Scan will read in a line typed at the console, and will
	// separate the values based on whitespace and put it in the memory addresses of the arguments.
	// The & symbol gets the memory address of a variable, same as C.
	fmt.Scan(&myName)
	// The Printf function is similar to the C function of the same name, it is useful if
	// you have data that you want to put in your printed text. Make sure you end with a
	// newline if you need one.
	fmt.Printf("It is nice to meet you, %s.\n", myName)
}
