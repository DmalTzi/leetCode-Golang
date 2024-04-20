package main

import (
    "fmt"
)

func main() {
    result := isPalindrome(121)
    fmt.Println(result)
}

func isPalindrome(x int) bool {
    var current int = x
    var reverse int = 0
    if x > 0 {
        for current != 0 {
            digit := current % 10
            reverse = reverse * 10 + digit
            current /= 10
            fmt.Println(reverse)
            if reverse == x {
                return true
            }
        }
    }


    return false
}
