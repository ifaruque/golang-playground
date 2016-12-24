package main

import "fmt"

func main() {
	var ptr *int
	var x = 100;

	ptr = &x;
	fmt.Println("pointer value", ptr)

	fmt.Println("Address of variable x = ", &x)
	fmt.Println("Value of pointer ptr and address of x is", ptr == &x)
	fmt.Println("Value of x using pointer = ", *ptr)

	var pptr **int

	pptr = &ptr
	fmt.Println("Value of pptr", **pptr)

	var ppptr ***int

	ppptr = &pptr
	fmt.Println("Value of ppptr", ***ppptr)

	***ppptr = 20;

	fmt.Println("Value of variable x = ", x)
}