package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(twoSum(nums, target))
	nums = []int{3, 3}
	target = 6
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	// make a map, where key is the number and value is the index
	n := make(map[int]int)
	// iterate through nums
	for idx, num := range nums {
		// if the difference between target and num is a map key, return the value of the key and the current index
		if i, ok := n[target-num]; ok {
			return []int{i, idx}
		}
		// otherwise, add the number to the map as key and its index as value
		n[num] = idx
	}
	// if no solution is found, return an empty slice
	return []int{}
}
