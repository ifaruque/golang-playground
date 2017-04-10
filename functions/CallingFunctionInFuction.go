package main

import "fmt"

func f1(param1, param2, param3 int) {
	fmt.Println("F1 is called with: ", param1, param2, param3)
}

func f2(a, b int) (rParam1, rParam2, rParam3 int) {
	fmt.Println("F2 is called")
	rParam1, rParam2 = b, a
	rParam3 = a + b
	return
}

func main() {
	f1(f2(2,3))
	fmt.Println("Done")
}