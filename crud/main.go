package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"net/http"
)

type Post struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
}

func main() {
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