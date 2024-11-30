package main

import (
	"fmt"
)

func ZipString(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		count := 1
		for i+1 < len(s) && s[i] == s[i+1] {
			count++
			i++ 
		}
		result += fmt.Sprintf("%d%c", count, s[i])
	}

	return result
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
