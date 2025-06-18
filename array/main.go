package main

import "fmt"

func main() {
	var name[5] string
	name[0] = "Rashi"
	name[1] = "umang"

	fmt.Println("Hello", name)

	var age = [5]int{21,43, 12, 34, 56} 
	fmt.Println("Age:", age)
	fmt.Println("Length of age array:", len(age)) // length of the array
	fmt.Println("name at index 1:", name[1]) // accessing the 2nd element of the array

	var price[5]float64 
	fmt.Println("Price:", price) // accessing the array of float64
	var user[5]string
	fmt.Printf("Username %q\n",user)   // %q(quoted string) is used to print the string in double quotes
}