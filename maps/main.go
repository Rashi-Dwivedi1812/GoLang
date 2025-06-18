package main

import "fmt"

func main() {
	studentGRades := make(map[string]int)
	studentGRades["Alice"] = 85
	studentGRades["Bob"] = 90
	studentGRades["Charlie"] = 78

	fmt.Println("Student Grades")
	for name, grade := range studentGRades {
		fmt.Printf("%s: %d\n", name, grade)
	}
	delete(studentGRades, "Bob")
	grades , exists := studentGRades["Bob"]
	fmt.Println("bob marks:", grades, "\nbob exists:", exists)


	person := map[string]int{
		"John": 30,
		"Jane": 25,	
		"Mike": 35,
}
	fmt.Println("Person Ages")
	for name, age := range person {
		fmt.Printf("%s: %d\n", name, age)
	}		
}