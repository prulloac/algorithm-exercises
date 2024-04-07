package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}

func isValid(s string) bool {
	// begin with a stack of runes
	stack := []rune{}
	// iterate through the string
	for _, r := range s {
		// if the rune is an opening parenthesis, add it to the stack
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
		} else {
			// if the stack is empty, return false
			if len(stack) == 0 {
				return false
			}
			// if the rune is a closing parenthesis, pop the last rune from the stack
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// if the rune is not a matching closing parenthesis, return false
			if (r == ')' && last != '(') || (r == ']' && last != '[') || (r == '}' && last != '{') {
				return false
			}
		}
	}
	// if the stack is empty, return true
	return len(stack) == 0
}
