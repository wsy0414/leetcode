package main

import "fmt"

func main() {
  fmt.Println(hasDuplicate([]int{1,2,3,3}))
  fmt.Println(hasDuplicate([]int{1,2,3,4}))
}

func hasDuplicate(nums []int) bool {
  m := make(map[int]int)
  for i, v := range nums {
    if _, ok := m[v]; ok {
      return true
    }
    m[v] = i
  }
  return false
}
