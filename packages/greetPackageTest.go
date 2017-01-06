package main

import (
	"fmt"
	"./greetings"
)

func main() {
	fmt.Println(greetings.GreetMe("Hello", "Nur Rony"))
	// fmt.Println(greetings.greet("Hello")) cannot refer to unexported name greetings.greet
}
