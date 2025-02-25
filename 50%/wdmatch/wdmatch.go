package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	string1, string2 := os.Args[1], os.Args[2]
	counter1, counter2 := 0, 0

	// Loop through s2 to find s1 in order
	for counter2 < len(string2) {
		if counter1 < len(string1) && string1[counter1] == string2[counter2] {
			counter1++
		}
		counter2++
	}

	// If we have matched all characters of s1 in s2, print s1
	if counter1 == len(string1) {
		fmt.Println(string1)
	}
	/*import "github.com/01-edu/z01" if fmt is not allowed
	if counter1 == len(string1) {
		for _, i := range string1 {
			z01.PrintRune(i)
		}
	}
	*/
}
