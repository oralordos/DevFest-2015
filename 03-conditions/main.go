package main

import "fmt"

func main() {
	// If you use :=, Go will create and fill the variable at the same time.
	x := 6 * 2
	// No parenthesis required, but the braces must start on the same line.
	if x < 8 {
		fmt.Println("This is a small number")
		// Note the else is on the same line as the else brace.
	} else if x > 17 {
		fmt.Println("This is a big number")
	} else {
		fmt.Println("This is a middle-sized number")
	}
}
