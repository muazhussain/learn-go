package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// first one inclusive and second one is exclusive
	b := a[:]   // all the elements
	c := a[3:]  // from 4th elements to the end
	d := a[:5]  // first 5 elements
	e := a[3:6] // 4th, 5th and 6th element
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
