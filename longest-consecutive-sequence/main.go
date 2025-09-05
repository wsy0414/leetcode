package main

import "fmt"

func main()  {
  fmt.Println(longestConsecutive([]int{2,20,4,10,3,4,5}))
  fmt.Println(longestConsecutive([]int{0,3,2,5,4,6,1,1}))
}

func longestConsecutive(nums []int) int {
  m := make(map[int]int, 0)

  for i, v := range nums {
    
    m[v] = i
  }

  result := 0
  for _, v := range nums {
    // 先確認這個數字是不是一個連續數列的開頭，用該數字-1來判斷，不是的話就跳過
    if _, ok := m[v-1]; ok {
      continue
    }
    cnt, curr := 0, v
    // 找到一個連續數列的起點後開始計算這個連續數列的長度
    for _, ok := m[curr]; ok; _, ok = m[curr] {
      cnt++
      curr++
    }
    // 如果長度超過目前結果就替換掉
    if cnt > result {
      result = cnt
    }
  }
  return result
}
