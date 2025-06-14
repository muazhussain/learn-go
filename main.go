package main

import "fmt"

func main() {
	statePopulation := map[string]int{
		"Dhaka":  150000,
		"Khulna": 2000,
	}
	fmt.Println(statePopulation)
	// Array can be key
	m := map[[3]int]int {}
	fmt.Println(m)

	// Slices can't be key
	n := map[[]int]int {}
	fmt.Println(n)
}
