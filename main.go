package main

import "fmt"

func main() {
	grades := [...]int{3, 4, 5}
	fmt.Printf("%v %T\n", grades, grades)
}
