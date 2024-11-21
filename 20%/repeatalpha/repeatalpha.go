package main

import (
	"fmt"
)

func RepeatAlpha(s string) string {

	var result []rune

	for _, i:=range(s){
		if i=='a' || i=='A'{
			for o:=0; o<1; o++{
				result = append(result,i)
		}
	}else if i=='b' || i=='B'{
		for o:=0; o<2; o++{
			result = append(result,i)
		}
	}else if i=='c' || i=='C'{
		for o:=0; o<3; o++{
			result = append(result,i)
		}
	}else if i=='d' || i=='D'{
		for o:=0; o<4; o++{
			result = append(result,i)
		}
	}else if i=='h' || i=='H'{
		for o:=0; o<9; o++{
			result = append(result,i)
		}
	}else if i=='o' || i=='O'{
		for o:=0; o<15; o++{
			result = append(result,i)
		}
	}else if i=='u' || i=='U'{
		for o:=0; o<18; o++{
			result = append(result,i)
		}
	}else if i=='m' || i=='M'{
		for o:=0; o<13; o++{
			result = append(result,i)
		}
	}else if i=='i' || i=='I'{
		for o:=0; o<10; o++{
			result = append(result,i)
		}
	}else{
		result = append(result,i)
	}
}
	
return string(result)
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
