package main

import ("os")
import ("github.com/01-edu/z01")

func main(){
	if len(os.Args)>4 || len(os.Args) < 4{
		return 
	}
		args:=os.Args[1:]
		var replaced string
		
		for _, i:=range args[0]{
			if string(i)==args[1]{
				replaced+=args[2]
			}else{
				replaced+=string(i)
			}
		}

		for _, i:=range replaced{
			z01.PrintRune(i)
		}
	z01.PrintRune('\n')
}