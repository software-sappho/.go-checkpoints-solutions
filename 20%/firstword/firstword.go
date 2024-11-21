package main

import (
	"fmt"
	"strings"
)

func FirstWord(s string) string {
	words := strings.Fields(s)
	res := "\n"
	if len(words) > 0 {
		res = words[0] + res
	}
	return res
}

func main() {
	fmt.Println(FirstWord("hello there"))
	fmt.Println(FirstWord(""))
	fmt.Println(FirstWord("hello   .........  bye"))
}
