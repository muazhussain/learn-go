package main

import "fmt"

func main() {
	// short declaration
	name := "Muaz"
	var isFunny = true

	fmt.Println(name, isFunny)

	// declare multiple variables in the same line
	firstName, lastName, age := "Muaz", "Hussain", 26

	fmt.Println(firstName, lastName, age)
}
