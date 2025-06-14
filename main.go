package main

import "fmt"

func main() {
	statePolulation := make(map[string]int)
	statePolulation = map[string]int{
		"Dhaka":  150000,
		"Khulna": 2000,
	}
	fmt.Println(statePolulation)
	statePolulation["Jashore"] = 1000
	fmt.Println(statePolulation)
}
