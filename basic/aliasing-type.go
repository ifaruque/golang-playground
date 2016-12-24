package main

import "fmt"

func main() {
	type TZ int
	type Rope string
	var rope Rope = "This is string alias"
	var a, b, c TZ = 1, 2, 0
	c = a + b

	fmt.Println("a =", a, "b =", b, "c =", c)
	fmt.Println("String alias", rope)
}
