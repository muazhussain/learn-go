package main

import "fmt"

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:     2,
		actorName:  "Muaz",
		companions: []string{"X", "Y", "Z"},
	}
	fmt.Println(aDoctor)
}
