package main

import (
	"fmt"
)

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	// If the number is negative, make it positive and note the sign
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	var result string

	// Build the string by taking the remainder (digits) of the number
	for n > 0 {
		result = string(n%10+'0') + result // Add the current digit (n % 10) at the front
		n /= 10                            // Move to the next digit
	}

	// If the number was negative, prepend the '-' sign
	if negative {
		result = "-" + result
	}

	return result
}
