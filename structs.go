package main

import ("fmt"
		"strconv")

// Define person struct
type Person struct{
	// firstName string
	// lastName string
	// city string
	// gender string
	// age int
	firstName, lastName, city, gender string
	age int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
  p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
	  p.lastName = spouseLastName
	
	}
}

func main() {
	// Init person using struct
	// person1 := Person{firstName: "Hana", lastName: "Tanaka", city: "Tokyo", gender: "f", age: 26}
	
	// Alternative
	person1 := Person{"Shin", "Tanaka", "Tokyo", "f", 26}
	
	person1.hasBirthday()
	person1.getMarried("williams")
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)
	fmt.Println(person1.greet())
}
