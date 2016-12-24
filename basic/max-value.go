package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	var ret int

	ret = max(a, b)

	fmt.Println("Max value is ", ret)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
