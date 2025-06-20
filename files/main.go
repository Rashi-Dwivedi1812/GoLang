package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("example.txt")      // Create a new file named output.txt
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	} 
	defer file.Close()

	content := "Hello, World!\nThis is a simple file writing example in Go.\n"
	_, errors := io.WriteString(file,content+"\n") // Write content to the file
	if errors != nil {		
		fmt.Println("Error writing to file:", errors)
		return
	}
}	