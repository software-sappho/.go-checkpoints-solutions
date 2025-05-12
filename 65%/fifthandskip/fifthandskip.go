package main

import "fmt"

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	// Count real (non-space) characters
	countNonSpace := 0
	for _, ch := range str {
		if ch != ' ' {
			countNonSpace++
		}
	}

	if countNonSpace < 5 {
		return "Invalid Input\n"
	}

	// Build a cleaned version without spaces
	var clean []rune
	for _, ch := range str {
		if ch != ' ' {
			clean = append(clean, ch)
		}
	}

	var result []rune
	i := 0
	for i+5 <= len(clean) {
		result = append(result, clean[i:i+5]...) // Add 5 characters
		result = append(result, ' ')             // Add space
		i += 6                                   // Skip 6th char
	}

	// Add any remaining characters (less than 5 left, no skip needed)
	for ; i < len(clean); i++ {
		result = append(result, clean[i])
	}

	// Remove trailing space if it ends with space
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	result = append(result, '\n')
	return string(result)
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
