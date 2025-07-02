package main

import (
	"fmt"
)

func main() {
  fmt.Println(isAnagram("racecar", "carrace"))
  fmt.Println(isAnagram("jar", "jam"))
}

// 使用map儲存出現的字符次數，最後比對兩個map內的資料是否都一樣
func isAnagram(s string, t string) bool {
  if len(s) != len(t) {
    return false
  }
  // 這是處理非英文時計算長度的方式，這樣才會準確
  /**if utf8.RuneCountInString(s) != utf8.RuneCountInString(t) {
    return false
  }*/

  // type rune = int32 主要處理中文和特殊符號utf-8
  m1, m2 := make(map[rune]int), make(map[rune]int)
  for i, word := range s {
    m1[word]++
    m2[rune(t[i])]++
  }

  for k, v := range m1 {
    if v != m2[k] {
      return false
    }
  }
  
  return true
}

// This solution assumes the input strings contain only lowercase English letters.
func isAnagramO1Space(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    // Create a fixed-size array to store character counts.
    // The size is 26, for 'a' through 'z'.
    // This is constant space, O(1), because it doesn't grow with input size.
    counts := [26]int{}

    // Iterate through the strings, incrementing for s and decrementing for t.
    for i := 0; i < len(s); i++ {
        counts[s[i]-'a']++
        counts[t[i]-'a']--
    }

    // If the strings are anagrams, all counts in the array will be zero.
    for _, count := range counts {
        if count != 0 {
            return false
        }
    }

    return true
}