package main

import (
	"gofr.dev/pkg/gofr"
	"fmt"
	// "myLearning/myUtil"
)
func main() {
	// fmt.Println("Hello, World!")

	// myutil.PrintMsg("This is a message from myUtil package.")


	// // Variable declaration and initialization
	// var name string = "Rashi"
	// fmt.Printf("Hello, %s!\n", name)
	// var age int = 25
	// fmt.Printf("You are %d years old.\n", age)
	// var height float64 = 5.5
	// height = 5.5 // This is a shorthand declaration
	// fmt.Printf("Your height is %.1f feet.\n", height)
	// var isStudent bool = true
	// if isStudent {
	// 	fmt.Println("You are a student.")
	// } else {
	// 	fmt.Println("You are not a student.")
	// }

	// var person = "I am Cute"
	// fmt.Println(person)
	// const pi float64 = 3.14159
	// fmt.Printf("The value of pi is approximately %.5f.\n", pi);

	// user := "Rashi"
	// fmt.Printf("Hello, %s!\n", user)

	// // initializing variables with capital letters to use them outside the package
	// // and private variables with lowercase letters to restrict access within the package
	// var PublicVar string = "This is a public variable"
	// fmt.Println(PublicVar)

	// var privateVar string = "This is a private variable"
	// fmt.Println(privateVar)

	app := gofr.New()

	app.GET("/",RootHandler)
	app.GET("/hello",HelloHandler)
	app.GET("/sql",SqlHandler)
	app.Run()
}

func RootHandler(ctx *gofr.Context)(any,error){
	return "Welcome to gofr",nil
}

func HelloHandler(ctx *gofr.Context)(any,error){
	return "hello",nil
}

func SqlHandler(ctx *gofr.Context) (any, error) {
	var name string
	err := ctx.SQL.QueryRow("SELECT name FROM users WHERE id = ?", 10001).Scan(&name)
	if err != nil {
		return nil, fmt.Errorf("DB error: %w", err)
	}
	return map[string]string{"message": fmt.Sprintf("Hello %s!", name)}, nil
}


//docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30