package main

import "fmt"

func ifNegative(a []string, n int) int {
	if n < 0 {
		n = len(a) + n
	}

	if n < 0 {
		n = 0
	} else if n > len(a) {
		n = len(a)
	}

	return n
}

func Slice(a []string, nbr ...int) []string {
	if len(nbr) == 0 {
		return a
	}

	first := nbr[0]
	if len(nbr) == 1 {
		if first < 0 {
			first = len(a) + first
			if first < 0 {
				return a
			}
		}
		return a[first:]
	}
	second := nbr[1]

	first = ifNegative(a, first)
	second = ifNegative(a, second)

	if first > second {
		return nil
	}

	return a[first:second]
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}
