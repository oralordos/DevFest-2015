package main

import (
	"fmt"
	"sync"
	"time"
)

const numThreads = 5

// The wait group type is designed to let you wait until other goroutines are finished.
var wg sync.WaitGroup

// There is nothing special about this function.
func printing(message string, waitTime time.Duration) {
	time.Sleep(waitTime)
	fmt.Println(message)
	// We need to tell the wait group that we are done.
	wg.Done()
}

func main() {
	for i := 0; i < numThreads; i++ {
		// We need to tell the wait group that we are about to have another goroutine running.
		wg.Add(1)
		// We call the function like we normally do, but we prefix the call with go.
		// This will perform the function at the same time as we move on with the rest of the code.
		go printing(fmt.Sprintf("Message %d", i), time.Duration(numThreads-i)*time.Second)
	}
	// Normally, all n messages would have to be printed before we get here.
	fmt.Println("Done looping")
	// When the main function ends, all goroutines also end, so we need to make sure we wait until we are done.
	wg.Wait()
}
