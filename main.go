package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a, 10)
	fmt.Printf("Capacity: %v\n", cap(a))
}
