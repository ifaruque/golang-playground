package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 30
	size := unsafe.Sizeof(i)

	fmt.Println("Size of integer in my computer", size)
}
