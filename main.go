package main

import "fmt"

func main() {
	statePolulation := make(map[string]int)
	statePolulation = map[string]int{
		"Dhaka":  150000,
		"Khulna": 2000,
	}
	// Passed as reference
	sp := statePolulation
	fmt.Println(statePolulation)
	fmt.Println(sp)
	delete(sp, "Dhaka")
	fmt.Println(statePolulation)
	fmt.Println(sp)
}
