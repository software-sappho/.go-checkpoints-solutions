package main

import (
	"fmt"
)

func PrintIfNot(str string) string {
	var returned string

	if len(str) < 3 {
		returned = "G\n"
	} else {
		returned = "Invalid Input\n"
	}

	return returned
}

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}
