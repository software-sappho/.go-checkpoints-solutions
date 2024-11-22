package main

import (
	"fmt"
)

func FishAndChips(n int) string {
var returned string

if n <0{
	returned="error: number is negative"
}else if (n%2)==0 && (n%3)==0{
	returned="fish and chips"
}else if (n%3)==0{
	returned="chips"
}else if (n%2)==0{
	returned="fish"
}else{
	returned="error: non divisible"
}
return returned
}


func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}
