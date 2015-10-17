package main

import "fmt"

// If you have a value already, Go will figure out the type by itself.
var story = "Once upon a time, %s %s a %s and then %s %s\n"

func main() {
	// You can define multiple variables on a line like this.
	var noun, verb, noun2, verb2 string
	fmt.Print("Enter a noun and a verb: ")
	// Scan can take more than one argument at a time, it will wait until it reads
	// enough to fill all the arguments.
	fmt.Scan(&noun, &verb)
	fmt.Print("Enter an animal: ")
	fmt.Scan(&noun2)
	fmt.Print("Enter a past tense verb: ")
	fmt.Scan(&verb2)
	fmt.Printf(story, noun, verb, noun2, noun, verb2)
}
