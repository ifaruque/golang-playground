package main

import "fmt"

func init() {
	fmt.Println("Calling init function");
}

func main() {
	fmt.Println("Calling main function")
	var str, result = "asSASA ddd dsjkdsjs dk", make(map[string]int)
	for _, ch := range str {
		result[string(ch)]++
	}


	for ch, val := range result {
		fmt.Println(ch + "= ", val);
	}
}
