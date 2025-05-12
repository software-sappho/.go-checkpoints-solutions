package main

import (
	"fmt"
	"os"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

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

func main() {
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}

	numStr := os.Args[1]
	num, valid := stringToInt(numStr)
	if !valid || num <= 0 {
		fmt.Println(0)
		return
	}

	sum := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
