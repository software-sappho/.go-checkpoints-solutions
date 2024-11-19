package main

import (
	"fmt"
)

func RetainFirstHalf(str string) string {
	var result string

	if len(str) == 0 {
		result = ""
	} else if len(str) == 1 {
		result = str
	} else {
		midpoint := len(str) / 2
		result = str[:midpoint]
	}

	return string(result)
}

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}
