package main

import "fmt"

const (
	_ = (20 * iota)
	a
	b
	c
)

func main() {
	fmt.Printf("%v %v %v\n", a, b, c)
}
