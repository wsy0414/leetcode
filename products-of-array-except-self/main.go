package main

import (
	"fmt"
)

func main() {
  fmt.Println(productExceptSelf([]int{1,2,4,6}))
  fmt.Println(productExceptSelf([]int{-1,0,1,2,3}))
  fmt.Println(productExceptSelfOpt([]int{1,2,4,6}))
  fmt.Println(productExceptSelfOpt([]int{-1,0,1,2,3}))
}

// 先計算每個數字前面的乘積，在計算後面的乘積，最後一起計算
func productExceptSelf(nums []int) []int {
  pres := make([]int, len(nums))
  sufs := make([]int, len(nums))
  prefix := 1
  for i := 0; i < len(nums); i++ {
    pres[i] = prefix
    prefix = prefix * nums[i]
  }
  suffix := 1
  for i := len(nums) - 1; i >= 0; i-- {
    sufs[i] = suffix
    suffix = suffix * nums[i]
  }

  var result []int
  for i := 0; i < len(pres); i++ {
    result = append(result, pres[i] * sufs[i])
  }

  return result
}

// this method space complexity is O(1)
func productExceptSelfOpt(nums []int) []int {
  result := make([]int, len(nums))
  prefix := 1
  for i := 0; i < len(nums); i++ {
    result[i] = prefix
    prefix = prefix * nums[i]
  }
  suffix := 1
  for i := len(nums) - 1; i >= 0; i-- {
    result[i] = result[i] * suffix
    suffix = suffix * nums[i]
  }

  return result
}
