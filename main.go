package main

import "fmt"

func main() {
	statePopulation := map[string]int{
		"Dhaka":   1000,
		"Khulna":  250,
		"Jashore": 120,
	}
	if pop, ok := statePopulation["Jashore"]; ok {
		fmt.Println(pop)
	} else {
		fmt.Println("Info does not exist!")
	}
}
