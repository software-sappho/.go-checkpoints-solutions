package main

import (
	"fmt"
)

func IsCapitalized(s string) bool {
	var ifi bool
	for _, i := range s {
		if (i == ' ') && !(i+1 >= 'A' || i+1 <= 'Z') {
			ifi = true
		} else {
			ifi = false
		}
	}

	return ifi
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
