package main

import (
    "fmt"
)

func main() {
    result := anwser([]int{3,2,4}, 6)
    fmt.Println(result)
}
func anwser(nums []int, target int) []int {
    dataSet := make(map[int]int)
    fmt.Println("dataset : ",dataSet)
    for i,v := range nums{
        fmt.Println("i and v :",i, v)
        value, ok := dataSet[target - v]
        fmt.Println("dataSet", dataSet[target - v])
        fmt.Println("value and ok", value, ok)
        if ok{
            fmt.Println(value, ok)
            return []int{i,value}
        }
        dataSet[v] = i
    }
    return []int{}
}

func twoSum(nums []int, target int) []int {
    var result []int
    fmt.Println(len(nums), target)
    for  i := 0; i < len(nums); i++ {
        for j := len(nums) - 1; j > i; j-- {
            if nums[i] + nums[j] == target {
                fmt.Printf("%v + %v = %v ", nums[i], nums[j], nums[i] + nums[j])
                result = append(result, i, j)
                return result
            }
        }
        fmt.Println("")
    }
    return result
}

