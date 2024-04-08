package main

import "fmt"

func main() {
	fmt.Println(true == isAnagram("anagram", "nagaram"))
	fmt.Println(false == isAnagram("rat", "car"))
	fmt.Println(false == isAnagram("aacc", "ccac"))
}

func isAnagram(a string, b string) bool {
	chars := make([]int, 26)
	chars = strKeyInc(a, chars)
	chars = strKeyDec(b, chars)
	for _, c := range chars {
		if c != 0 {
			return false
		}
	}
	return true
}

func strKeyInc(s string, chars []int) []int {
	for i := 0; i < len(s); i++ {
		chars[s[i]-'a']++
	}
	return chars
}

func strKeyDec(s string, chars []int) []int {
	for i := 0; i < len(s); i++ {
		chars[s[i]-'a']--
	}
	return chars
}
