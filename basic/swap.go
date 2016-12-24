package main

import "fmt"

func main() {
	var x, y int = 2, 3
	fmt.Println("Before swap x = ", x, " y = ", y)

	x, y = swap(x, y)
	fmt.Println("Before swap x = ", x, " y = ", y)

}
func swap(a, b int) (int, int) {
	return b, a
}
