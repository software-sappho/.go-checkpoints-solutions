package main

import (
    "fmt"
   
)

func LastWord(s string) string {
    lastWord := ""
    tempWord := ""
    
    for i := 0; i < len(s); i++ {
        if s[i] != ' ' { // Build the current word
            tempWord += string(s[i])
        } else if tempWord != "" { // Word ends when a space is encountered
            lastWord = tempWord // Update lastWord with the current word
            tempWord = ""       // Reset tempWord
        }
    }
    
    // If the last character(s) is a word (not spaces), set lastWord to tempWord
    if tempWord != "" {
        lastWord = tempWord
    }
    
    return lastWord + "\n"
}


func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}

