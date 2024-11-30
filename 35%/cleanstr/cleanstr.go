package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Print("\n")
		return
	} else {

		args := os.Args[1:]
		returned := ""

		for i := 0; i < len(args[0]); i++ {

			returned += string(args[0][i])

			if i > 0 && args[0][i] == ' ' && args[0][i-1] == ' ' {
				returned = returned[:len(returned)-1]
			}

		}

		fmt.Println(returned)
	}
}
