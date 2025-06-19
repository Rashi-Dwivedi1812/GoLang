package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple, banana, cherry"
	parts := strings.Split(data, ",")
	fmt.Println("Parts of the string:", parts)

	str := "one two three four two two five"
	count := strings.Count(str, "two")
	fmt.Println("Count of 'two' in the string:", count)

	str = "      Hello, World!    "
	trimmed := strings.TrimSpace(str)
	fmt.Println("Trimmed string:", trimmed)

	str1 := "Hello"
	str2 := "world"
	result := strings.Join([]string{str1, str2}, ",")
	fmt.Println("Joined string:", result)
}