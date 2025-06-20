package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")                                               //1
	defer fmt.Println("This is a simple Go program.")                          //4
	defer fmt.Println("It demonstrates basic output functionality.")           //3
	fmt.Println("You can run this code to see the output in your terminal.")   //2
}