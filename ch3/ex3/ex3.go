package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	// Struct literal style without names
	firstEmployee := Employee{"Paco", "Fernandez", 0}

	// Struct literal style with names
	secondEmployee := Employee{
		firstName: "Eufrasio",
		lastName:  "Rodriguez",
		id:        1,
	}

	// Var declaration
	var thirdEmployee Employee
	thirdEmployee.firstName = "Pepe"
	thirdEmployee.lastName = "Muñoz"
	thirdEmployee.id = 3

	fmt.Println(firstEmployee)
	fmt.Println(secondEmployee)
	fmt.Println(thirdEmployee)

}
