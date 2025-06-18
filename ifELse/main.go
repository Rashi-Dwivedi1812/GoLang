package main

import "fmt"

func main() {
	// if-else statement
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}

	// if-else if-else statement
	y := 20
	if y < 10 {
		fmt.Println("y is less than 10")
	} else if y < 30 {
		fmt.Println("y is between 10 and 30")
	} else {
		fmt.Println("y is greater than or equal to 30")
	}

	z := 15
	if z < 10 && x > 5 {
		fmt.Println("z is less than 10")
	} else {
		fmt.Println("z is greater than or equal to 20")
	}
}