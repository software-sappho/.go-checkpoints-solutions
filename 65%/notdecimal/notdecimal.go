package main

import "fmt"

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	// Check if the string contains a decimal point
	hasDot := false
	var afterDot []rune
	var beforeDot []rune

	// Split the string at the decimal point
	for i, ch := range dec {
		if ch == '.' {
			hasDot = true
			beforeDot = []rune(dec[:i])
			afterDot = []rune(dec[i+1:])
			break
		}
	}

	// If no decimal point, return the number as is
	if !hasDot {
		return dec + "\n"
	}

	// Special case: if the string starts with 0. something, remove the 0
	if len(beforeDot) == 1 && beforeDot[0] == '0' && len(afterDot) > 0 && afterDot[0] != '0' {
		beforeDot = []rune{}
	}

	// Remove leading zeros from the part after the decimal point
	for len(afterDot) > 0 && afterDot[0] == '0' {
		afterDot = afterDot[1:]
	}

	// If the afterDot part is empty after removing leading zeros, just return beforeDot
	if len(afterDot) == 0 {
		return string(beforeDot) + "\n"
	}

	// Check if the string is a valid number
	for i, ch := range dec {
		if ch == '-' && i == 0 {
			continue
		}
		if ch == '.' {
			continue
		}
		if ch < '0' || ch > '9' {
			return dec + "\n"
		}
	}

	// Combine beforeDot and afterDot parts and return the result
	var result []rune
	result = append(result, beforeDot...)
	result = append(result, afterDot...)
	result = append(result, '\n')
	return string(result)
}

func main() {
	fmt.Print(NotDecimal("0.1"))       // Should output 1
	fmt.Print(NotDecimal("174.2"))     // Should output 1742
	fmt.Print(NotDecimal("0.1255"))    // Should output 1255
	fmt.Print(NotDecimal("1.20525856")) // Should output 120525856
	fmt.Print(NotDecimal("-0.0f00d00")) // Should output -0.0f00d00
	fmt.Print(NotDecimal(""))          // Should output newline
	fmt.Print(NotDecimal("-19.525856")) // Should output -19525856
	fmt.Print(NotDecimal("1952"))      // Should output 1952
}
