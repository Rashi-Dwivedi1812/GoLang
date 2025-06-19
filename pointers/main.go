package main

import "fmt"

func modifyValuebyReference(value *int) {
	*value = 100
}

func main() {
	//Method 1:
	var num int = 10
	var ptr *int = &num // Pointer to num, which holds the address of num, 
	// *int is a pointer type describing the type of the variable it points to
	
	//Method 2: 
	num1 := 20
	ptr1 := &num1
	
	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", ptr)
	fmt.Println("Value of num:", num1)
	fmt.Println("Address of num:", *ptr1)


	var pointer *int
	if pointer == nil {	
		fmt.Println("Pointer is nil")
	} else {
		fmt.Println("Pointer is not nil")
	}
	fmt.Println("Value of pointer:", pointer) // This will print <nil> since pointer is not initialized
	fmt.Println("Address of pointer:", &pointer) // This will print the address of the pointer


	value := 42
	modifyValuebyReference(&value) // Passing the address of value
	fmt.Println("Value after function call:", value)
}
