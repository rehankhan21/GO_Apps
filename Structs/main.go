package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}

func main() {
	// relies on the order of the fields for assignment
	// not ideal
	// alex := person{"Alex", "Anderson"}

	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// fmt.Println(alex)

	// automatically adds zero values not undefined 
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Printf("%+v", alex)

	jim := person {
		firstName: "Jim",
		lastName: "Party",
		contact: contactInfo{
			email: "jim@gmail.com",
			zipCode: 9768,
		},
	}

	fmt.Printf("%+v", jim)
}