package main

import "fmt"

func divide(a, b float64) (float64,error) {    //error is a built-in type in Go, used to represent errors
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b,nil
}

func main() {
	fmt.Println("started error handling")
	ans,_ := divide(10, 0)  // The second return value is an error, but we are ignoring it with the underscore 
	fmt.Println("Answer is:", ans)
}