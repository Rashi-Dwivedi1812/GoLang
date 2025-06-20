package main

import (
	"fmt"
	"net/url"
)

func main() {
	myURL := "https://jsonplaceholder.typicode.com/posts/1?key=1" // Define the URL to fetch data from

	parsedURL, err := url.Parse(myURL) // Parse the URL string into a URL object
	if err != nil {
		fmt.Println("Error parsing URL:", err) 
		return
	}
	fmt.Printf("Type of parsed URL: %T\n", parsedURL) 

	fmt.Println("Scheme:", parsedURL.Scheme) 
	fmt.Println("Host:", parsedURL.Host) 
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Raw Query:", parsedURL.RawQuery)

	//Modify the URL components
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "key=2" 
	
	newURL := parsedURL.String() // Convert the modified URL object back to a string
	fmt.Println("Modified URL:", newURL) // Print the modified URL 
}