package main

import (
	"fmt"
	// "io"
	"io/ioutil"
	// "os"
)

func main() {
	// file, err := os.Create("example.txt")      // Create a new file named output.txt
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	return
	// } 
	// defer file.Close()

	// content := "Hello, World!\nThis is a simple file writing example in Go.\n"
	// byte, errors := io.WriteString(file,content+"\n") // Write string to the file
	// fmt.Println("Wrote", byte, "bytes to the file.")
	// if errors != nil {		
	// 	fmt.Println("Error writing to file:", errors)
	// 	return
	// }

	// file, err := os.Open("example.txt") // Open an existing file named example.txt
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// defer file.Close()

	// buffer := make([]byte, 1024) // Create a buffer to hold file content

	// for{
	// 	n,err := file.Read(buffer) // Read from the file into the buffer
	// 	if err == io.EOF{          /// Check for end of file
	// 		break;
	// 	}
	// 	if err != nil {
	// 		fmt.Println("Error reading file:", err)
	// 		return
	// 	}
	// 	fmt.Print(string(buffer[:n])) // Print the content read from the file
	// }

	// Read the entire file content into memory 
	content,err := ioutil.ReadFile("example.txt") // Read the entire file content into memory
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:\n", string(content)) // Print the file content as a string
}	