package main

import "fmt"

func main() {
	i1 := 5
	fmt.Printf("An integer : %d, it's location in memory: %p\n", i1, &i1)

	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

	// String pointer
	var str string = "Good bye"
	var pStr *string = &str

	*pStr = "Caio"

	fmt.Printf("Here is the pointer: %p\n", pStr)
	fmt.Printf("Here is the string accessed with pointer: %s\n", *pStr)
	fmt.Printf("Here is the string accessed directly: %s\n", str)
}
