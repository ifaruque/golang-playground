package main

import (
	"./greetings"
	"fmt"
	myfmt "./fmt"
)

func main() {
	fmt.Println(greetings.GreetMe("Hello", "Nur Rony"))
	myfmt.PrintNewln("Hello","world")
	// fmt.Println(greetings.greet("Hello")) cannot refer to unexported name greetings.greet
}
