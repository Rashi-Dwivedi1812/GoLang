package main

import "fmt"

func main() {
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")	
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}

	month := "January"
	switch month {
	case "January","February", "March":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Spring")		
		default:
		fmt.Println("Summer or Autumn")
}

	temperature := 30
	switch {
	case temperature < 0:	
		fmt.Println("Freezing")
	case temperature>=0 && temperature < 20:	
		fmt.Println("Cold")
	case temperature>=20 && temperature < 30:	
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
}
}