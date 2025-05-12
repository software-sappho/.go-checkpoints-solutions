package main

import (
	"fmt"
	"os"
)

// Function to convert string to int manually
func stringToInt(s string) (int, bool) {
	result := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, false
		}
		result = result*10 + int(c-'0')
	}
	return result, true
}

func fprime(value int) {
	if value == 1 || value == 0 {
		return
	}
	divisionIterator := 2
	for value > 1 {
		if value%divisionIterator == 0 {
			fmt.Print(divisionIterator)
			value = value / divisionIterator

			if value > 1 {
				fmt.Print("*")
			}
			divisionIterator--
		}
		divisionIterator++
	}
	fmt.Println()
}

func main() {
	if len(os.Args) == 2 {
		numStr := os.Args[1]
		num, valid := stringToInt(numStr)
		if valid {
			fprime(num)
		}
	}
}
