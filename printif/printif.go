package main

import (
	"fmt"
)

func PrintIf(str string) string {
	var returned string

	if len(str) >= 3{
		returned = "G\n"
	}else if len(str)==0{
		returned = "G\n"
	} else {
		returned = "Invalid Input\n"
	}

	return returned
}

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}
