package main

import (
	"fmt"
	"os"
)

func toUpper(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - ('a' - 'A')
	}
	return b
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

func main() {
	for _, arg := range os.Args[1:] {
		argBytes := []byte(arg)
		for i := 0; i < len(argBytes); i++ {
			if i+1 == len(argBytes) || argBytes[i+1] == ' ' {
				argBytes[i] = toUpper(argBytes[i])
			} else {
				argBytes[i] = toLower(argBytes[i])
			}
		}
		fmt.Println(string(argBytes))
	}
}
