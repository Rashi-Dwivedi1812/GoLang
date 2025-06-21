package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	person := Person{
		Name:    "John Doe",
		Age:     30,
		IsAdult: true,
	}
	fmt.Println("Person Struct:", person)

	//convert struct to JSON encoding (Marshalling)
	jsonData,err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println("JSON Data:", string(jsonData))

	//convert JSON encoding to struct (Unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println("Unmarshalled Person Struct:", personData)
}