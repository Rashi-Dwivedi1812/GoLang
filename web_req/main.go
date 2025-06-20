package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res,err := http.Get("http://jsonplaceholder.typicode.com/posts/1") // Make a GET request to the specified URL
	if err != nil {
		fmt.Println("Error making GET request:", err) // Print error if the request fails
		return
	}
	defer res.Body.Close() // Ensure the response body is closed after reading
	fmt.Printf("Type of response: %T\n", res)

	//Read the response body
	data,err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err) 
		return
	}
	fmt.Println("Response body:", string(data)) 

}