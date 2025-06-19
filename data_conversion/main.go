package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 10
	fmt.Println("Value of num:", num)
	fmt.Printf("Type of num:%T\n", num)

	var data float64 = float64(num)
	data += 5.5 
	fmt.Println("Value of data:", data)
	fmt.Printf("Type of data:%T\n", data)

	num = 123
	str := strconv.Itoa(num) // Convert int to string
	fmt.Println("Value of str:", str)
	fmt.Printf("Type of str:%T\n", str)

	number_string := "456"
	number, _ := strconv.Atoi(number_string) // Convert string to int
	fmt.Println("Value of number:", number)
	fmt.Printf("Type of number:%T\n", number)

	num_string := "789.12"
	num_float,_ := strconv.ParseFloat(num_string, 64) // Convert string to float
	fmt.Println("Value of num_float:", num_float)
	fmt.Printf("Type of num_float:%T\n", num_float)
}