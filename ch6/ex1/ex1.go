package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func main() {
	p := MakePerson("Fred", "Williamson", 25)
	fmt.Println(p)
	p2 := MakePersonPointer("Wilma", "Fredson", 32)
	fmt.Println(p2)
}
