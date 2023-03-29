package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// relies on the order of the fields for assignment
	// not ideal
	// alex := person{"Alex", "Anderson"}

	alex := person{firstName: "Alex", lastName: "Anderson"}

	fmt.Println(alex)
}