package main

import (
	"fmt"
	"sync"
)

func workers(i int,wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter
	fmt.Printf("Worker %d started\n", i)
	//some tasks is happening 
	fmt.Printf("Worker %d completed its task\n", i)
}

func main() {
	var wg sync.WaitGroup
	//Start 3 workers goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go workers(i,&wg)
	}

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("Task completed")
}