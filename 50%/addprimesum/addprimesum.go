package main

import (
    "fmt"
    "os"
    "strconv"
)

func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println(0)
        return
    }

    num, err := strconv.Atoi(os.Args[1])
    if err != nil || num <= 0 {
        fmt.Println(0)
        return
    }

    sum := 0
    for i := 2; i <= num; i++ {
        if isPrime(i) {
            sum += i
        }
    }
    fmt.Println(sum)
}
