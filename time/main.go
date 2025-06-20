package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)
	fmt.Printf("type: %T\n", currentTime)

	formattedTime := currentTime.Format("2006-01-02,Monday,03:04 PM")
	fmt.Println("Formatted Time:", formattedTime)

	layoutstr := "2006-01-02,03:04 PM"
	dateStr := "2023-10-01,12:00 PM"
	parsedTime, _ := time.Parse(layoutstr, dateStr)
	fmt.Println("Parsed Time:", parsedTime)

	//Add 1 more day to the current time
	oneDayLater := currentTime.Add(24 * time.Hour)
	fmt.Println("One Day Later:", oneDayLater)
	formatted := oneDayLater.Format("2006-01-02,Monday,03:04 PM")
	fmt.Println("Formatted One Day Later:", formatted)
}