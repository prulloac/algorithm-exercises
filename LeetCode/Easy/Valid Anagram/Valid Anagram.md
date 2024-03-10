# Valid Anagram

#HashTable, #String, #Easy, #Sorting

Given two strings `s` and `t`, return true if `t` is an anagram of `s`, and false otherwise.

> An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

## Example 1:

> Input: s = "anagram", t = "nagaram"
>
> Output: true

## Example 2:

> Input: s = "rat", t = "car"
>
> Output: false

## Constraints:

- `1 <= s.length, t.length <= 5 * 104`
- `s` and `t` consist of lowercase English letters.

# Source
https://leetcode.com/problems/valid-anagram/

# Solutions

## Naive

> Sort both strings and compare them, if they are equal, return true, otherwise return false.

- Time complexity: `O(n log n)`
- Space complexity: `O(1)`

## Not-so-naive

> Compare both strings length, if they are not equal, return false, then loop through the first string and compare character at index `i` with the character at index `length - i - 1` of the second string, if they are not equal, return false, otherwise return true.

- Time complexity: `O(n)`
- Space complexity: `O(1)`

## Fast-code Solution

> - Reverse one of the strings and compare them, if they are equal, return true, otherwise return false.

- Time complexity: `O(n)`
- Space complexity: `O(1)`
