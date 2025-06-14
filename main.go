package main

import "fmt"

func main() {
	statePolulation := make(map[string]int)
	statePolulation = map[string]int{
		"Dhaka":  150000,
		"Khulna": 2000,
	}
	pop, ok := statePolulation["Dhaka"]
	fmt.Println(pop, ok)
	pop, ok = statePolulation["Jashore"]
	fmt.Println(pop, ok)
}
