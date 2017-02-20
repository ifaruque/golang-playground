package main

import "fmt"

type innerStruct struct {
	innerInt1 int
	innerInt2 int
}

type outerStruct struct {
	outerInt1 int
	outerFloat float32
	int
	innerStruct
}

func main() {
	outerStructObject := new(outerStruct)
	outerStructObject.innerInt1 = 1
	outerStructObject.innerInt2 = 2
	outerStructObject.outerInt1 = 3
	outerStructObject.outerFloat = 4.0

	// we access anonymous struct field with it's datatype. So each struct
	// can have only one anonymous filed for each datatype.
	outerStructObject.int = 5
	fmt.Println("with `new` operator", outerStructObject, &outerStructObject)

	outerStructObject = &outerStruct{1, 2.0, 3, innerStruct{3, 4}}
	fmt.Println("with `&` operator", outerStructObject, &outerStructObject)

}
