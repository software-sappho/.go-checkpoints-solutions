package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	} else {
		args := os.Args[1:]
		returned := ""

		inWord := false

		for i := 0; i < len(args[0]); i++ {
			char := args[0][i]

			if char == ' ' {
				if inWord {
					returned += "   "
					inWord = false 
				}
			} else {
				if !inWord {
					inWord = true
				}
				returned += string(char)
			}
		}

		if len(returned) > 0 && returned[len(returned)-3:] == "   " {
			returned = returned[:len(returned)-3]
		}

		if len(returned) > 0 {
			fmt.Println(returned)
		}
	}
}
