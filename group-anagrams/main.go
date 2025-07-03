package main

import (
	"fmt"

)

func main() {
  fmt.Println(groupAnagrams([]string{"act", "post", "tops", "cat", "stop", "hat"}))
  fmt.Println(groupAnagrams([]string{"x"}))
  fmt.Println(groupAnagrams([]string{""}))
}

type cnt [26]int

/**
我只需要知道每個字串的英文字母出現的頻率，然後用這個頻率當作map的key去建立分類
*/
func groupAnagrams(strs []string) [][]string {
  m := make(map[cnt][]string)

  for _, s := range strs {
    var counts cnt
    for i := 0; i < len(s); i++ {
      counts[s[i] - 'a']++
    }
    
    m[counts] = append(m[counts], s)
  }

  result := [][]string{}
  for _, v := range m {
    result = append(result, v)
  }

  return result
}
