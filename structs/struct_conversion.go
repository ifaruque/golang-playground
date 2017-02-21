package main
import "fmt"
type number struct {
	f float32
}

type nr number   // alias type

func main() {
	a := number{5.0}
	b := nr{5.0}
	// var i float32 = b
	// var i = float32(b)
	// var c number = b
	var c = nr(b) // valid: number(b)
	fmt.Println(a, b, c)
}
