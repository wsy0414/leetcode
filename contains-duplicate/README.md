# Contains Duplicate

## Problem Description

Given an integer array `nums`, return `true` if any value appears at least twice in the array, and return `false` if every element is distinct.

**Example 1:**

*   **Input:** `nums = [1,2,3,1]`
*   **Output:** `true`

**Example 2:**

*   **Input:** `nums = [1,2,3,4]`
*   **Output:** `false`

**Example 3:**

*   **Input:** `nums = [1,1,1,3,3,4,3,2,4,2]`
*   **Output:** `true`

## Test Cases

*   `[]int{1, 2, 3, 3}` -> `true`
*   `[]int{1, 2, 3, 4}` -> `false`
*   `[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}` -> `true`

## My Anser Complexity Analysis

*   **Time Complexity:** O(n), where n is the number of elements in the array. This is because we iterate through the array once.
*   **Space Complexity:** O(n), where n is the number of elements in the array. In the worst case, we might have to store all the elements in the hash map.

