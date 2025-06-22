package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	// "io/ioutil"
	"net/http"
)

type Post struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
}

func performGetRequest() {
	res,err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {         // Check if the status code is OK (200)
		fmt.Println("Error: Status code is not OK:", res.Status)
		return
	
	}
	// data,err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return
	// }
	// fmt.Println("Response Data:", string(data))

	var post Post
	err = json.NewDecoder(res.Body).Decode(&post)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("Post Data:", post)
	fmt.Printf("Post ID: %d, Title: %s\n", post.ID, post.Title)
}

func performPostRequest() {	
	post := Post{
		UserID: 1,
		Title:  "Rashi",
		Body:   "Dwivedi",
	}

	// Convert the post struct to JSON
	jsonData, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	//convert the jsonData to a string
	jsonString := string(jsonData)

	//convert string to reader
	jsonReader := strings.NewReader(jsonString)

	myurl := "https://jsonplaceholder.typicode.com/posts"
	// Perform the POST request
	res,err := http.Post(myurl,"application/json",jsonReader)
	if err != nil {
		fmt.Println("Error performing POST request:", err)
		return	
	}
	defer res.Body.Close()

	data,_ := ioutil.ReadAll(res.Body)
	fmt.Println("Response Data:", string(data))

	fmt.Println("Response Status:", res.Status)
}

func performUpdateRequest() {
	post := Post{
		UserID: 18,
		Title:  "Rashi cutie",
		Body:   "hello world",
	}
	// Convert the post struct to JSON
	jsonData, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	//convert the jsonData to a string
	jsonString := string(jsonData)

	//convert the jsonData to a string
	jsonReader := strings.NewReader(jsonString)
	const myurl = "https://jsonplaceholder.typicode.com/posts/1"

	///create put request
	req,err := http.NewRequest(http.MethodPut, myurl, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing PUT request:", err)
		return
	} 
	defer res.Body.Close()
	data,_ := ioutil.ReadAll(res.Body)
	fmt.Println("Response Data:", string(data))
	fmt.Println("Response Status:", res.Status)
}

func performDeleteRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/posts/1"

	///create delete request
	req,err := http.NewRequest(http.MethodDelete, myurl,nil)
	if err != nil {
		fmt.Println("Error creating delete request:", err)
		return
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing PUT request:", err)
		return
	} 
	defer res.Body.Close()
	fmt.Println("Response Status:", res.Status)
}

func main() {
	// performGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	performDeleteRequest()
}