package main

import (
	"fmt"
	"math"
)

func main() {
	const pi float64 = math.Acos(-1)
	fmt.Printf("%v %T\n", pi, pi)
}
