package main

import "fmt"

func main() {
	a := make([]int, 3, 100) // type, length, capacity 
	fmt.Printf("len: %v, cap: %v\n", len(a), cap(a))
}
