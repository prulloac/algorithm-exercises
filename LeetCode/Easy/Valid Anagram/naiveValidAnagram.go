package main

import (
	"fmt"
	s "strings"
)

func main() {
	fmt.Println(true == isAnagram("anagram", "nagaram"))
	fmt.Println(false == isAnagram("rat", "car"))
	fmt.Println(false == isAnagram("aacc", "ccac"))
}

func isAnagram(a string, b string) bool {
	// to first check if length is the same, it's not return false
	if len(a) != len(b) {
		return false
	}
	// iterate through characters of first string
	for _, c := range a {
		// if character is not present or different quantity, return false
		C := string(c)
		if s.Index(b, C) < 0 || s.Count(a, C) != s.Count(b, C) {
			return false
		}
	}
	return true
}
