package main

import (
	"strconv"
	"fmt"
)

func main() {
	var orig string = "ABC"
	var an int
	var err error
	an, err = strconv.Atoi(orig)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("Orig %s is not an integer - exiting with error %v\n", orig, err)
		// os.Exit(1)
	}
	orig = "913"
	an, err = strconv.Atoi(orig)

	fmt.Printf("The integer is %d and error is %v\n", an, err)

	k := 6
	switch k {
		case 4: fmt.Println("was <= 4"); fallthrough;
		case 5: fmt.Println("was <= 5"); fallthrough;
		case 6: fmt.Println("was <= 6"); fallthrough;
		case 7: fmt.Println("was <= 7"); fallthrough;
		case 8: fmt.Println("was <= 8"); fallthrough;
		default: fmt.Println("default case")
	}
}