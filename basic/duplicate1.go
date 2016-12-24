package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}
	for length, number := range counts {
		if number > 1 {
			fmt.Printf("%d\t%s\n", number, length)
		}
	}
}
