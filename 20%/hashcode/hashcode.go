package main

import (
	"fmt"
)

func HashCode(dec string) string {

	len :=len(dec)
	var hex string

	for _, i:=range dec{
		hexxed:=(int(i)+len)%127
		if hexxed < 33 || hexxed >127{
			hexxed += 33
		}
		hex += string(hexxed)
	}
	return hex
}


func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
