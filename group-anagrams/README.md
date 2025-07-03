# Group Anagrams

## Problem Description

Given an array of strings `strs`, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

## Test Cases

Example 1:
```go
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
```

Example 2:
```go
Input: strs = [""]
Output: [[""]]
```

Example 3:
```go
Input: strs = ["a"]
Output: [["a"]]
```

## My Answer Complexity Analysis

- **Time Complexity:** O(n*k) where n is the number of strings and k is the maximum length of a string.
- **Space Complexity:** O(n*k)

