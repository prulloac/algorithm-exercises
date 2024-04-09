package main

import (
	"fmt"
)

func main() {
	expected := []int{1, 2}
	result := topKFrequentElements([]int{1, 1, 1, 2, 2, 3}, 2)
	fmt.Println(compareSlices(expected, result))
}

func topKFrequentElements(nums []int, k int) []int {
	// create a hash table
	table := make(map[int]int)
	for _, i := range nums {
		table[i]++
	}
	// create a reverse index table
	revtable := make(map[int][]int)
	for n, count := range table {
		// key = frequency, value = numbers with same frequency
		revtable[count] = append(revtable[count], n)
	}
	// create result list
	list := make([]int, 0)
	// reverse iteration
	for i := len(nums); len(list) != k; i-- {
		for _, n := range revtable[i] {
			if len(list) != k {
				list = append(list, n)
			}
		}
	}
	return list
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
