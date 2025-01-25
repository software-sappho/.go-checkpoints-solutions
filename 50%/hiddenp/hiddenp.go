package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var i int
	if len(os.Args) != 3 {
		return
	}
	for _, r := range os.Args[1] {
		j := strings.Index(os.Args[2][i:], string(r))
		if j == -1 {
			fmt.Println("0")
			return
		}
		i += j + 1
	}
	fmt.Println("1")
}

// second solution

/*

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	source := os.Args[1]
	target := os.Args[2]
	targetIndex := 0

	for i := 0; i < len(source); i++ {
		found := false
		for j := targetIndex; j < len(target); j++ {
			if source[i] == target[j] {
				targetIndex = j + 1
				found = true
				break
			}
		}
		if !found {
			fmt.Println("0")
			return
		}
	}
	fmt.Println("1")
}


*/