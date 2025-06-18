package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your name:")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("hello", name)

	reader := bufio.NewReader(os.Stdin)  // Create a new buffered reader to read input from standard input
	user, _ := reader.ReadString('\n')
	fmt.Println("Hello", user)
}