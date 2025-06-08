package main

import "fmt"

const (
	a = iota
	b
	c
)

func main() {
	fmt.Printf("%v %v %v\n", a, b, c)
}
