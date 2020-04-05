package main

import "fmt"

//struct to represent a person
//firstname :string and lastname:string

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex.firstName, alex.lastName)
	var alex person
	alex.firstName = "alex"
	alex.lastName = "Anderson"
	fmt.Println(alex) //{alex Anderson}
	fmt.Printf("%v%v", alex.firstName, alex.lastName)
	fmt.Printf("%+v", alex) //{firstName:alex lastName:Anderson}
}
