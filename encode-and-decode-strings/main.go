package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
  s1 := []string{"neet","code","love","you"}
  s2 := []string{"we","say",":","yes"}
  
  s3 := encode(s1)
  s4 := encode(s2)
  fmt.Println(s3)
  fmt.Println(s4)
  fmt.Println(decode(s3))
  fmt.Println(decode(s4))
}

// 把每個字串的長度放在字串前面中間用符號區隔，之後要拆解時先取長度再根據長度擷取文字
func encode(strs []string) string {
  var sb strings.Builder
  for _, s := range strs {
    sb.WriteString(strconv.Itoa(len(s)))
    sb.WriteString(":")
    sb.WriteString(s)
  }

  return sb.String()
}

func decode(str string) []string {
  idx, i := 0, 0
  var result []string
  for i < len(str) {
    if str[i] != ':' {
      i++
      continue
    }
    strLen, err := strconv.ParseInt(str[idx:i], 10, 0)
    if err != nil {
      return []string{}
    }
    result = append(result, str[i+1:i+1+int(strLen)])
    i += (int(strLen) + 1)
    idx = i
  }

  return result
}
