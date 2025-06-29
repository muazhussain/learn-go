package main

import "fmt"

func main() {
	fmt.Println("first")
	// defer execute at the end just before the function return
	defer fmt.Println("second")
	fmt.Println("third")
}
