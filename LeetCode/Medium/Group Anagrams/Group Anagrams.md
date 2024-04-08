# Group Anagrams

#Array, #HashTable, #String, #Medium, #Sorting

Given an aray of strings `strs`, group **the anagrams** together. You can retrn the answer in **any order**.

> An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

## Example 1:

> Input: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]\
> Output: [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]

## Example 2:

> Input: strs = [""]\
> Output: [[""]]

## Example 3:

> Input: strs = ["a"]\
> Output: [["a"]]

## Constraints:

- `1 <= strs.length, 10^4`
- `0 <= strs[i].length <= 100`
- `strs[i]` consists of lowercase English letters.

# Source
https://leetcode.com/problems/group-anagrams/

# Solutions

## Naive

> Create a hash table and loop through the strings, make a copy of the string, sort it, and use it as a key in the hash table, if the key is not in the hash table, add it with an empty list as a value, then append the original string to the list. Finally, return the values of the hash table as a list.

- Time complexity: `O(n * m log m)`
- Space complexity: `O(n * m)`
