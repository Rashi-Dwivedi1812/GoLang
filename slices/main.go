package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}     //Method 1: Using a slice literal
	fmt.Println("Numbers:", numbers)
	numbers = append(numbers, 6, 7, 8)
	fmt.Println("Updated Numbers:", numbers)
	fmt.Printf("Data Type of numbers: %T\n", numbers)
	fmt.Println("Length of numbers:", len(numbers))

	names := []string{}     //Method 2: Using a slice literal with an empty slice
	fmt.Println("Names:", names)

	items := make([]string, 3, 5)      //Method 3: Using the make function
	items = append(items, "apple")
	items = append(items, "banana")
	items = append(items, "cherry")  //it will double the capacity 

	fmt.Println("Items:", items)
	fmt.Println("Length of items:", len(items))
	fmt.Println("Capacity of items:", cap(items))
}