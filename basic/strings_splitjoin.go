package main

import (
	"strings"
	"fmt"
)

func main() {
	str := "A quick brown fox jumps over the lazy dog"

	fmt.Println("Main string =", str)
	/**
	  strings.Fields(s) splits the string s around each instance of one or more consecutive white space characters,
	  returning a slice of substrings []string of s or an empty list if s contains only white space
	 */
	sl := strings.Fields(str)

	fmt.Printf("Splitted in slice: %v\n", sl)

	for _, s := range sl {
		fmt.Printf("%s - ", s)
	}

	fmt.Println()

	// strings.Split(str, sep)
	str = "GO1|The ABC of Go|25"
	sp := strings.Split(str, "|");

	fmt.Printf("After strings.Slice = ", sp)
	fmt.Println()

	fmt.Println("After strings.Join", strings.Join(sl, "-"))
}
