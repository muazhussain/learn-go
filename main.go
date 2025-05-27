package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Type conversion
	var i int = 42
	var k float32
	k = float32(i)
	fmt.Printf("%v %T\n", k, k)

	var pi float32 = 3.1416
	var absPI int = int(pi)
	fmt.Printf("%v %T\n", absPI, absPI)

	var j int = 10
	var val string = strconv.Itoa(j)
	fmt.Printf("%v %T\n", val, val)
}
