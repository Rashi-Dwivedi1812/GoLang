package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email   string
	Phone   string
	Address string
}

type Address struct {
	House int
	City    string
	State   string
}

type Employee struct {
	Person_details Person
	Contact_details Contact
	Address_details Address
}

func main() {
	// Method 1: Using a struct literal
	var Person1 Person
	fmt.Println("Name:",Person1)
	Person1.FirstName = "John"
	Person1.LastName = "Doe"
	Person1.Age = 30
	fmt.Println("Name:", Person1.FirstName, Person1.LastName, "Age:", Person1.Age) 

	// Method 2
	Person2 := Person{
		FirstName: "Jane",
		LastName:  "Smith",
		Age:       25,
	}
	fmt.Println("Person 2:", Person2)

	// Method 3: Using new to create a pointer to a Person struct
	var Person3 = new(Person)
	Person3.FirstName = "Alice"
	Person3.LastName = "Johnson"	
	Person3.Age = 28
	fmt.Println("Person 3:", Person3)

	var employee1 Employee
	employee1.Person_details = Person{
		FirstName: "Bob",
		LastName:  "Brown",
		Age:       35,
	}
	employee1.Contact_details.Email = "contact@gmail.com"
	employee1.Contact_details.Phone = "123-456-7890"
	employee1.Address_details = Address{
		House: 123,
		City:    "New York",
		State:   "NY",
	}
	fmt.Println("Employee1 Details:", employee1)
}