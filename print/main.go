package main

import "fmt"

func main() {
	age := 25
	name := "Rashi"
	height := 5.5

	fmt.Printf("Hello, World!")
	fmt.Printf("Name: %s, Age: %d, Height: %.1f\n", name, age, height)
	
	fmt.Println("age:", age, "name:", name, "height:", height)

	fmt.Printf("type of age: %T", age) //tyope of age: int
}