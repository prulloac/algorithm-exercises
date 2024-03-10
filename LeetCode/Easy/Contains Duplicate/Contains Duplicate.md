# Contains Duplicate

#Array, #HashTable, #Easy, #Sorting

Given an integer array `nums`, return true if any value appears at least twice in the array, and return false if every element is distinct. 

## Example 1:

> Input: nums = [1,2,3,1]
>
> Output: true

## Example 2:

> Input: nums = [1,2,3,4]
>
> Output: false

## Example 3:

> Input: nums = [1,1,1,3,3,4,3,2,4,2]
>
> Output: true

# Constraints:

- 1 <= nums.length <= 105
- -109 <= nums[i] <= 109

# Source
https://leetcode.com/problems/contains-duplicate/

# Solutions

## Naive

> Check if any element in the array is duplicated by comparing each element with the rest of the array.

- Time complexity: `O(n^2)`
- Space complexity: `O(1)`

## Not-so-naive

> Instantiate a set and loop through the array, adding each element to the set. If the element is already in the set, return true. If the loop finishes, return false.

- Time complexity: `O(n)`
- Space complexity: `O(n)`

## Fast-code Solution

> Given that set does not have duplicates, instantiate a set based on the array and check if length of the set is different from the length of the array.

- Time complexity: `O(n)`
- Space complexity: `O(n)`
