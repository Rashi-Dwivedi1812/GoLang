package main

import (
	"fmt"
	"time"
)

func sayHello() { 
	fmt.Println("Hello, World!")
	// time.Sleep(2 * time.Second) // Simulate a delay
	fmt.Println("Hello again after a short delay!")
}

func sayGoodbye() {
	fmt.Println("Goodbye, World!")
	time.Sleep(1 * time.Second) // Simulate a delay
	fmt.Println("Goodbye again after a short delay!")
}

func main() {
	fmt.Println("Starting the program...")
	go sayHello()
	go sayGoodbye()

	time.Sleep(800 * time.Millisecond) // Wait for a while before exiting
}