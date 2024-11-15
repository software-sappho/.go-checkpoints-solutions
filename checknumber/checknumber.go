package main

import (
	"fmt"
)

func CheckNumber(arg string) bool {
	var isnumber bool

	for _, i := range arg {
		if i == '1' || i == '2' || i == '3' || i == '4' || i == '5' || i == '6' || i == '7' || i == '8' || i == '9' || i == '0' {
			isnumber = true
		} else {
			isnumber = false
		}
	}
	return isnumber
}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}
