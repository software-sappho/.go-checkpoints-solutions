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
	first, second := 0, 0

	// Loop through s2 to find s1 in order
	for second < len(string2) {
		if first < len(string1) && string1[first] == string2[second] {
			first++
		}
		second++
	}

	// If we have matched all characters of s1 in s2, print s1
	if first == len(string1) {
		fmt.Println(string1)
	}
}
