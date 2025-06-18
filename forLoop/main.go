package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Iteration:", i)
	// }

	// count := 0
	// for {
	// 	fmt.Println("Infinite loop")
	// 	count++
	// 	if count >= 5 {
	// 		break
	// 	}
	// }

	num := []int{1, 2, 3, 4, 5}
	for index,value := range num{
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	data := "Hello, World!"
	for index, char := range data{
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}