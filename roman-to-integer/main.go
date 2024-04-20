package main

import (
    "fmt"
)

func main() {
    result := romanToInt("MCMXCIV")
    fmt.Println(result)
}

func romanToInt(s string) int {
    var number int // declare number 
    var romanSymbol map[string]int = map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
        "IV": -2,
        "IX": -2,
        "XL": -20,
        "XC": -20,
        "CD": -200,
        "CM": -200,
    } // delceare romanSymbol for map
    for i := 0; i < len(s); i++ {
        number += romanSymbol[string(s[i])] // map in romanSymbol and plus
        if i+2 <= len(s) {
            dbC := s[i:i+2] // get Double Char from s
            number += romanSymbol[dbC] //map dbC in romanSymbol and plus with negative number
        }
    }
    return number
}
