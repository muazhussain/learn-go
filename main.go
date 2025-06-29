package main

import "fmt"

func main() {
	// last defer executes first (Last in First out)
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
}
