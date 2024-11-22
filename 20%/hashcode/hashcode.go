package main

import (
	"fmt"
)

func HashCode(dec string) string {

	size:=len(dec)
	var hashed string

	for _, i:=range(dec){
		hash:=(int(i)+size)%127
		if hash <32 || hash > 127{
			hash+=33
		}
		hashed+=string(hash)
	}
	
	return hashed
}


func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
