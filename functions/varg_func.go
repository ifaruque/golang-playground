package main

import "fmt"

func main() {
	x := min(1, 2, 3, 4, 5)
	fmt.Println("Minimum number is ", x)

	arr := []int{6, 7, 8, 9}
	x = min(arr...)
	fmt.Println("Minimum number is ", x)

	unknownTypeParams(1, 3, "Hello", true);
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


func unknownTypeParams(params ...interface{})  {
	for _, value := range params {
		switch value.(type) {
		case int:
			fmt.Printf("%v type is %T \n", value, value)
		case string:
			fmt.Printf("%v type is %T \n", value, value)
		case bool:
			fmt.Printf("%v type is %T \n", value, value)

		}
	}
}