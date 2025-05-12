package main

import "fmt"

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Input\n"
	}

	// Remove spaces manually
	var noSpaces []rune
	for _, char := range str {
		if char != ' ' {
			noSpaces = append(noSpaces, char)
		}
	}

	// Build the result
	var result []rune
	count := 0
	for _, char := range noSpaces {
		if count == 5 {
			result = append(result, ' ')
			count = 0
		}
		result = append(result, char)
		count++
	}
	result = append(result, '\n')

	return string(result)
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
