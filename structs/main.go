package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	anna := person{firstName: "anna", lastName: "peterson", age: 30}
	fmt.Println(anna)
}
