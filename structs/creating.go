package main

import (
	"fmt"
	"unsafe"
)

type struct1 struct {
	i1 int
	f1 float32
	str string
}

func main() {
	ms := new(struct1)
	ms.i1 = 13
	ms.f1 = 13.13
	ms.str = "string value with new"

	fmt.Printf("the string is %s\n", ms.str)
	fmt.Printf("the integer is %d\n", ms.i1)
	fmt.Printf("the float is %f\n", ms.f1)
	fmt.Printf("size of staticMS %f\n", unsafe.Sizeof(ms))

	// With shorthand initialization

	ms2 := &struct1{13, 13.13, "string value with shorthand pointer"}

	fmt.Printf("the string is %s\n", ms2.str)
	fmt.Printf("the integer is %d\n", ms2.i1)
	fmt.Printf("the float is %f\n", ms2.f1)
	fmt.Printf("size of staticMS %f\n", unsafe.Sizeof(ms2))


	var staticMS struct1;
	staticMS.i1 = 13
	staticMS.f1 = 13.13
	staticMS.str = "Simple regular variable"

	fmt.Printf("the string is %s\n", staticMS.str)
	fmt.Printf("the integer is %d\n", staticMS.i1)
	fmt.Printf("the float is %f\n", staticMS.f1)
	fmt.Printf("size of staticMS %f\n", unsafe.Sizeof(staticMS))
}
