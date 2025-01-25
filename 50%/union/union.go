package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 3 {
		var res string
		s1 := os.Args[1]
		s2 := os.Args[2]

		for _, v := range s1 {
			if !strings.ContainsRune(res, v) {
				res += string(v)
			}
		}
		for _, v := range s2 {
			if !strings.ContainsRune(res, v) {
				res += string(v)
			}
		}
		fmt.Print(res)
	}
	fmt.Println()
}

// second solution

/*

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 3 {
		s1 := os.Args[1]
		s2 := os.Args[2]
		res := ""

		// Helper function to check if a character exists in a string
		contains := func(str string, char rune) bool {
			for _, v := range str {
				if v == char {
					return true
				}
			}
			return false
		}

		// Process the first string
		for _, v := range s1 {
			if !contains(res, v) {
				res += string(v)
			}
		}

		// Process the second string
		for _, v := range s2 {
			if !contains(res, v) {
				res += string(v)
			}
		}

		fmt.Print(res)
	}
	fmt.Println()
}


*/