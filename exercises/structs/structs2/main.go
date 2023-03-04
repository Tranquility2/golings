package main

import "fmt"

type Person struct {
	name string
	age  int
}

type ContactDetails struct {
	person Person
	phone  string
}

func main() {
	person := Person{name: "John", age: 32}
	contactDetails := ContactDetails{person: person, phone: "+01 333 666"}

	fmt.Printf("%s is %d years old and his phone is %s\n",
		contactDetails.person.name,
		contactDetails.person.age,
		contactDetails.phone)
}
