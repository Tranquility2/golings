package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	person := Person{name: "Dan", age: 21}
	fmt.Printf("Person %s and age %d", person.name, person.age)
}
