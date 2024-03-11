# Valid Parentheses

#String #Stack #Easy

Given a string s containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

## Example 1:

> Input: s = "()" \
> Output: true

## Example 2:

> Input: s = "()[]{}" \
> Output: true

## Example 3:

> Input: s = "(]" \
> Output: false

## Constraints:

- `1 <= s.length <= 10^4`
- `s` consists of parentheses only `'()[]{}'`.

# Source

https://leetcode.com/problems/valid-parentheses/

# Solutions

> Loop through the string and for each character, check if it is an open bracket, if so, push it to the stack, otherwise, check if the stack is empty, if so, return `false`, otherwise, pop the top element from the stack and check if it is the corresponding open bracket for the current character, if not, return `false`, otherwise, continue. After the loop, check if the stack is empty, if so, return `true`, otherwise, return `false`.

- Time complexity: `O(n)`
- Space complexity: `O(n)`

```python
