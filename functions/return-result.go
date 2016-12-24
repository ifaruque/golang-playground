package main

import "fmt"

func main() {
	var x, y, num1, num2 int = 2, 3, 0, 0

	passByRef(&x)
	fmt.Println("In main x =", x)

	num1, num2 = unnamedReturnFunc(x, y)
	fmt.Println("Unnamed Return: Number1 =", num1, "Number2 =", num2)

	num1, num2 = namedReturnFunc(num1, num2)
	fmt.Println("Named return: Number1 =", num1, "Number2 =", num2)

	add, mul, sub := AddMultiplySubNamed(num1, num2)
	fmt.Println("Named return: add =", add, "mul =", mul, "sub =", sub)

	add, mul, sub = AddMultiplySubUnnamed(add, sub)
	fmt.Println("Unamed return: add =", add, "mul =", mul, "sub =", sub)
}

func unnamedReturnFunc(x, y int) (int, int) {
	return x + 2, y + 2;
}

func namedReturnFunc(x, y int) (a int, b int) {
	a, b = x * 2, y * 2;
	return
}


func passByRef(b *int) {
	*b = 10
	fmt.Println("In pass by ref =",  *b)
}

func AddMultiplySubNamed(a, b int) (int, int, int) {
	return a + b, a * b ,a - b
}

func AddMultiplySubUnnamed(a, b int) (add, mul, sub int) {
	add = a + b
	mul = a * b
	sub = a - b

	return
}

