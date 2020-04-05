package main

import (
	"fmt"
)

//struct to represent a person
//firstname :string and lastname:string
//embedding contact struct

type contactInfo struct {
	email   string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@test.com",
			zipcode: 99889,
		},
	}
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
