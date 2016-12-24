package main

import (
	"runtime"
	"fmt"
)

func main() {
	fmt.Println("Program is running on ", runtime.GOOS)
}
