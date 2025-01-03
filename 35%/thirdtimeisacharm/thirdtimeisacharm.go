package main

import (
	"fmt"
)

func ThirdTimeIsACharm(str string) string {

	returned := ""

	if len(str)==0 || len(str)<3{
	 	return "\n"
	}

	for i:=1; i<len(str);i++{
		if (i%3)==2{
			returned+=string(str[i])
		}
	}
	returned+= "\n"
	return returned
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}
