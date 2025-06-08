package main

import "fmt"

func main() {
	a := [3]int{1, 3, 5}
	// copy (does not reference)
	b := &a
	b[1] = 4
	fmt.Println(a)
	fmt.Println(b)
}
