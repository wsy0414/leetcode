# Valid Anagram

## Problem Description

Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`, and `false` otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

## Test Cases

Example 1:
```go
Input: s = "racecar", t = "carrace"
Output: true
```

Example 2:
```go
Input: s = "jar", t = "jam"
Output: false
```

## My Answer Complexity Analysis

### Initial Solution (Hash Map)
- **Time Complexity:** O(n)
- **Space Complexity:** O(n)

This solution uses hash maps to count character frequencies and works with any character set.

### O(1) Space Solution (Array)
- **Time Complexity:** O(n)
- **Space Complexity:** O(1)

This improved solution uses a fixed-size array (size 26) to count character frequencies, assuming the input strings only contain lowercase English letters. This removes the dependency on input size for space.

## Note
Golang的string預設是UTF-8編碼的而且**for...range**在針對string時會使用utf-8套件內的方法確認文字是否為unicode，所以返回的型別會是rune,底層邏輯是runtime時會判斷byte的前幾位確認是不是unicode

