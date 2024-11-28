package main

import (
	"fmt"
)

func RepeatAlpha(s string) string {

	repeated:=[]rune{}

	for _, i:=range s{
		if i == 'a'|| i=='A'{
			for o:=0;o<1;o++{
			repeated=append(repeated,i)
			}
		}else if i == 'b'|| i=='B'{
			for o:=0;o<2;o++{
			repeated=append(repeated,i)
			}
		}else if i == 'c'|| i=='C'{
			for o:=0;o<3;o++{
			repeated=append(repeated,i)
			}
		}else if i == 'd'|| i=='D'{
			for o:=0;o<4;o++{
			repeated=append(repeated,i)
			}
		}else{
			repeated=append(repeated,i)
		}
	}

	return string(repeated)
}
	



func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
