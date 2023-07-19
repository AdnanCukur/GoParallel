package goparallel

import (
	"sync"
)

// Process processes a slice in parallel using specified maximum parallelism.
// The slice parameter is the input slice to be processed.
// It takes a processFunction that defines the processing logic for each item in the slice.
// The maxParallelism parameter specifies the maximum number of goroutines running in parallel.
func Process[T interface{}](slice []*T, processFunction func(*T), maxParallelism int) {
	var wg sync.WaitGroup
	wg.Add(len(slice))                         // Increment the WaitGroup counter.
	limiter := make(chan bool, maxParallelism) // Create a buffered channel to limit the number of parallel goroutines.

	for _, item := range slice {
		limiter <- true                                     // Add a value to the limiter channel, which blocks if the maximum parallelism is reached.
		go runFunction(processFunction, item, limiter, &wg) // Launch a goroutine to run the processFunction for the current item.
	}

	wg.Wait() // Wait for all goroutines to finish processing.
}

// runFunction runs the processFunction for a single input item.
// It takes the processFunction, the input item, the limiter channel, and the WaitGroup.
func runFunction[T interface{}](processFunction func(*T), input *T, limiter chan bool, wg *sync.WaitGroup) {
	defer wg.Done()        // Signal the WaitGroup that this goroutine is done.
	processFunction(input) // Execute the processFunction for the input item.
	<-limiter              // Remove a value from the limiter channel, allowing another goroutine to start.
}
