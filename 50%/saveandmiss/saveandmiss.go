package main

import (
	"fmt"
)

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}
	result := ""
	length := len(arg)
	for i := 0; i < length; i += 2 * num {
		// Save the next `num` characters
		end := i + num
		if end > length {
			end = length
		}
		result += arg[i:end]
	}
	return result
}

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}
