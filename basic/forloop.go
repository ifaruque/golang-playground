package main

import "fmt"

func main() {
	str := "Go is a beautiful language!"
	fmt.Println("String length", len(str))

	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c\n", ix, str[ix])
	}

	// With Unicode
	/**
	 * Note:
	 * normal ASCII-characters using 1 byte, an indexed character is the full character,
	 * whereas for non-ASCII characters (who need 2 to 4 bytes) the indexed character is no longer correct
	 */
	str2 := "রনী"
	fmt.Println("\nWith Unicode")
	fmt.Println("String length with unicode", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c\n", ix, str2[ix])
	}

	fmt.Println()

	for i := 1; i < 5; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}

	fmt.Println()
	for i, s := 1, "G"; i < 5; i, s = i + 1, s + "G" {
		fmt.Println(s)
	}

	fmt.Println("\nFIZBUZZ Problem")
	const (
		FIZZ = 3
		BUZZ = 5
		FIZZBUZZ = 15
	)

	for i := 1; i < 101; i++ {
		switch {
		case i % FIZZBUZZ == 0:
			fmt.Print ("FizzBuzz ")
		case i % FIZZ == 0:
			fmt.Print ("Fizz ")
		case i % BUZZ == 0:
			fmt.Print("Buzz ")
		// just for pretty output
		// can't choose 3/5 because switch in Go lang does not
		// fall through implicitly
		case i % 8 == 0:
			fmt.Println(i)
		default:
			fmt.Print(i, " ")
		}
	}

}