package main

import (
	"fmt"
	"os"
	"strings"
)

func result(s1 string, s2 string) string {
	var rest []rune
	for _, a := range s1 {
		for _, b := range s2 {
			if a == b && !strings.ContainsRune(string(rest), a) {
				rest = append(rest, a)
			}
		}
	}
	return string(rest)
}

func main() {
	if len(os.Args) == 3 {
		fmt.Println(result(os.Args[1], os.Args[2]))
	}
}

// second solution

/*

package main

import (
	"fmt"
	"os"
)

func result(s1 string, s2 string) string {
	var rest []rune
	for _, a := range s1 {
		found := false
		for _, r := range rest {
			if a == r {
				found = true
				break
			}
		}
		if found {
			continue
		}
		for _, b := range s2 {
			if a == b {
				rest = append(rest, a)
				break
			}
		}
	}
	return string(rest)
}

func main() {
	if len(os.Args) == 3 {
		fmt.Println(result(os.Args[1], os.Args[2]))
	} else {
		fmt.Println("Usage: <program> <string1> <string2>")
	}
}


*/