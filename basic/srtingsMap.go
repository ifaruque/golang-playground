package main

import (
	"fmt"
	"strings"
)

func main() {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+2)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+2)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))

	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitle("хлеб"))
	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n"))
}
