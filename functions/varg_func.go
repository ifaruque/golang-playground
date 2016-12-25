package main

import "fmt"

func main() {
	x := min(1, 2, 3, 4, 5)
	fmt.Println("Minimum number is ", x)

	arr := []int{6, 7, 8, 9}
	x = min(arr...)
	fmt.Println("Minimum number is ", x)
}

func min(a ...int) int {
	if len(a) == 0 {
		return 0
	}

	min := a[0]

	for _, v := range a {
		if v < min {
			min = v
		}
	}

	return min
}
