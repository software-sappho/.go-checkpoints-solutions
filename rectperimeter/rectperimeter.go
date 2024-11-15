package main

import (
	"fmt"
)

func RectPerimeter(w, h int) int {
	var result int

	if w <= 0 || h <= 0 {
		result = -1
	} else {
		result = 2*w + 2*h
	}

	return result
}

func main() {
	fmt.Println(RectPerimeter(10, 2))
	fmt.Println(RectPerimeter(434343, 898989))
	fmt.Println(RectPerimeter(10, -2))
}
