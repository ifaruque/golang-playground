package main

import "fmt"

type Family struct {
	name string
	age int
	email string
}

func main() {
	var Nur Family
	var Afifa Family

	Nur.name = "Nur Mohammed Rony"
	Nur.age = 31
	Nur.email = "pro.nmrony@yahoo.com"

	Afifa.name = "Afifa Nur"
	Afifa.age = 28
	Afifa.email = "afifanur2013@gmail.com"

	printFamily(&Nur)
	printFamily(&Afifa)
}
func printFamily(family *Family) {
	fmt.Println("Full name", family.name)
	fmt.Println("Age", family.age)
	fmt.Println("email", family.email)
}
