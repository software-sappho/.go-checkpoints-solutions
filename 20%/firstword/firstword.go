package main

import (
	"fmt"
)

func FirstWord(s string) string {
	// sliced := []rune
	first := []rune{}

	for _, i := range s {
		if i != ' ' {
			first = append(first, i)
		} else {
			break
		}
	}

	return string(first)
}

func main() {
	fmt.Println(FirstWord("hello there"))
	fmt.Println(FirstWord(""))
	fmt.Println(FirstWord("hello   .........  bye"))
}
