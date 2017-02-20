package main

import (
	"strings"
	"fmt"
)

type Person struct {
	firstName string
	lastName string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}
func main() {
	person := new(Person)
	person.firstName = "Nur"
	person.lastName = "Rony"
	fmt.Println("Before upPerson", person)
	upPerson(person)
	fmt.Println("after upPerson", person.firstName, person.lastName)
}
