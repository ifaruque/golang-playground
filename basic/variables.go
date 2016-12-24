package main

import "fmt"

func main() {
	var a, b, c = 3, 4, "foo"
	fmt.Println("Value of a is = ", a)
	fmt.Printf("Type of a is %T", a)

	fmt.Println("Value of b is = ", b)
	fmt.Printf("Type of b is %T", b)

	fmt.Println("Value of c is = ", c)
	fmt.Printf("Type of c is %T", c)

}
