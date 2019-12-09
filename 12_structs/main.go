package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName, lastName, city, gender string
	// lastName  string
	// city      string
	// gender    string
	age int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver because we are changing something)
func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// initialize Person
	person1 := Person{firstName: "Grace", lastName: "Lee", city: "Raleigh", gender: "f", age: 25}

	// shorter way
	person2 := Person{"John", "Kim", "Boston", "m", 35}

	// fmt.Println(person1, person2)
	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1.age)
	person1.hasBirthday()
	person1.getMarried("Park")
	person2.getMarried("Park")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
