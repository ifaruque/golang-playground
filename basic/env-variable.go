package main

import (
	"os"
	"fmt"
)

func main() {
	var (
		USER = os.Getenv("USER")
		GOPATH = os.Getenv("GOPATH")
		SHELL = os.Getenv("SHELL")
	)

	fmt.Println("User is", USER)
	fmt.Println("GOPATH", GOPATH)
	fmt.Println("SHELL", SHELL)

}
