package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
  fmt.Println(validPalindrome("Was it a car or a cat I saw?"))
  fmt.Println(validPalindrome("tab a cat"))
}

func validPalindrome(s string) bool {
  reg := regexp.MustCompile(`[^A-Za-z0-9]`)

	// 把不需要的字元替換成空字串
	s = reg.ReplaceAllString(s, "")
  s = strings.TrimSpace(s)
  s = strings.ToLower(s)
  ss := strings.Split(s, "")
  fmt.Println(s)
  for i := 0; i < len(ss) / 2; i++ {
    if ss[i] != ss [len(ss) - i - 1] {
      return false
    }
  }
  return true
}
