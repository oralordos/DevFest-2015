package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// When using random numbers, make sure to seed the generator
	rand.Seed(time.Now().UnixNano())

	// This is an array, it has a size
	numbers := [3]float64{3.14, 6.21, 543.21}
	// This is a slice, it is like the vector from C++, it uses an array internally, but
	// can grow and shrink. It also passes by reference, so it is faster to send into a function.
	data := []int{1, 2, 3, 4, 5, 6}
	// This a map, you can go from nearly any type to any type with this.
	favorites := map[string]int{
		"Golang": 42,
		"Random": rand.Intn(100),
		"Foo":    465,
		"Bar":    678,
	}
	fmt.Println(numbers)
	fmt.Println(data)
	fmt.Println(favorites)
	fmt.Println(favorites["Foo"])
	fmt.Println(data[0])

	x := numbers[1:]
	fmt.Println(x)
}
