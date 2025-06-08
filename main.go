package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Printf("slice: %v\nlength: %v\ntype: %T\n", a, len(a), a)
}
