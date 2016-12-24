package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v", s, s)
	}

	fmt.Println()

	v1 := int64(-42)

	s10 := strconv.FormatInt(v1, 10)
	fmt.Printf("%T, %v\n", s10, s10)

	s16 := strconv.FormatInt(v1, 16)
	fmt.Printf("%T, %v\n", s16, s16)

}