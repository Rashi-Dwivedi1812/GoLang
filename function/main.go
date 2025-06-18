package main

import "fmt"

func simplefunction() {
	fmt.Println("This is a simple function")
}

func add(a , b int) (result int) {     //b int means a,b both are an integer, result int means the function returns an integer
	result = a + b
	return
}

func main() {
	fmt.Println("hello world") 
	simplefunction() // Call the simple function

	sum := add(5, 10) // Call the add function
	fmt.Println("Sum of 5 and 10 is:", sum) // Print the result of the add function
}