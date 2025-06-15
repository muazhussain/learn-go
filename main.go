package main

import "fmt"

func main() {
	i := 16
	switch {
	case i <= 10:
		fmt.Println("Less than or equal to 10")
		fallthrough
	case i >= 15:
		fmt.Println("Less than or equal to 15")
		fallthrough
	case i <= 20:
		fmt.Println("Less than or equal to 20")
		fallthrough
	default:
		fmt.Println("Unknown")
	}
}
