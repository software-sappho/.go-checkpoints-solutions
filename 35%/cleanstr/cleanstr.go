package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 || len(os.Args[1]) == 0 {
		fmt.Print("\n")
		return
	}

	input := os.Args[1]
	var result string
	inWord := false

	for i := 0; i < len(input); i++ {
		if input[i] == ' ' || input[i] == '\t' {
			if inWord {
				result += " "
				inWord = false
			}
		} else {
			result += string(input[i])
			inWord = true
		}
	}

	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	fmt.Println(result)
}
