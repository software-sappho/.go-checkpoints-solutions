package main

import (
	"fmt"
)

func DigitLen(n, base int) int {
	var count int

	if base < 2 && base > 32 {
		count = -1
	} else {
		if n < 0 {
			// newn:=n*-1
			// for i:=0, i<=base,i++{
			// 	result:=newm/base
			// 	count++
			// 	if result ==0{
			// 	break
			// 	}
			count = 0
		} else {
			var i int
			for i = 0; i < 20; i++ {
				i++
				if base+n >= 20 {
					count = i
					break

				}

			}

		}
	}

	return count
}

func main() {
	fmt.Println(DigitLen(1, 1))
	fmt.Println(DigitLen(5, 2))
	fmt.Println(DigitLen(5, 1))
	fmt.Println(DigitLen(3, 8))
}
