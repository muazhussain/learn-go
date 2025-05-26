package main

import "fmt"

func main() {
	name := "Muaz"
	// Prints a formatted string in standard out 
	fmt.Printf("Hello, %v\n", name)

	age := 26
	// Returns a formatted string 
	formattedText := fmt.Sprintf("You are %v years old!\n", age)
	fmt.Println(formattedText)
}
