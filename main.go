package main

import "fmt"

func main() {
	// variable declaration using var keyword
	var smsSendingLimit int
	var costPerSMS float32
	var hasPermission bool
	var userName string

	fmt.Printf(
		"%d %f %t %s\n",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		userName,
	)
}