package main

import (
    "fmt"
)

func FirstWord(s string) string {
    firstWord := ""
    for i := 0; i < len(s); i++ {
        if s[i] != ' ' { // Add non-space characters to firstWord
            firstWord += string(s[i])
        } else if firstWord != "" { // Stop adding when a space is found after starting the word
            break
        }
    }
    return firstWord + "\n"
}

func main() {
    fmt.Print(FirstWord("hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}

