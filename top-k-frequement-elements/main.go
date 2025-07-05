package main

import "fmt"

func main()  {
  fmt.Println(topKFrequement([]int{1,2,2,3,3,3}, 2))  
  fmt.Println(topKFrequement([]int{1,2,2,3,3,}, 2))  
  fmt.Println(topKFrequement([]int{7,7}, 1))  
}

func topKFrequement(nums []int, k int) []int {
  count := make(map[int]int)
  freqs := make([][]int, len(nums)+1)
  
  // 算出每個數字總共出現幾次
  for _, num := range nums {
    count[num] ++
  }
  // 根據出現的次數放入陣列中相對應的位置
  for k, v := range count {
    freqs[v] = append(freqs[v], k)
  }

  var result []int
  // freqs陣列中在越後面的數字束線的次數越多，所以從後面開始塞入result
  for i := len(freqs) - 1; i > 0; i-- {
    for _, v := range freqs[i] {
      result = append(result, v)
      if len(result) == k {
        return result
      }
    }
  }

  return result
}
