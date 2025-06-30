package main

import "fmt"

func main() {
  fmt.Println(twoSum([]int{3,4,5,6}, 7))
  fmt.Println(twoSum([]int{4,5,6}, 10))
  fmt.Println(twoSum([]int{5,5}, 10))
}

func twoSum(nums []int, target int) []int {
  m := make(map[int]int, 0)

  for i, v := range nums {
    if idx, ok := m[target - v]; ok {
      return []int{idx, i}
    }
    m[v] = i
  }

  return []int{}
}
