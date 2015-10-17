package main

import "fmt"

// Recursive definition of fibonacci's sequence, it is a slow calculation this way.
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	// Channels must be made before they can be used.
	c := make(chan int)
	// I am making an anonymous function that runs fib(40), and running it in another goroutine.
	// The <- after a channel passes the data to the other end of the channel, waiting if there is nothing on the other end.
	// The return value of the fib call is being passed into the channel.
	go func() {
		c <- fib(40)
	}()
	// I can do something else while the goroutine runs.
	fmt.Println(fib(30))
	// The <- before a channel returns the data in the channel, waiting if there is no data.
	fmt.Println(<-c)
}
