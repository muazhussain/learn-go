package main

import "fmt"

func main() {
	s := "This is a string"
	// String is an array
	// These are aliases of bytes

	fmt.Printf("%T %v\n", s[2], s[2])

	// String are imutable
	// This is an error
	s[2] = "u"
}