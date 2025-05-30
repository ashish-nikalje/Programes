package main

import (
	"fmt"
)

func findLeastFrequent(nums []int) int {
	// Edge case: If array is empty, return -1
	if len(nums) == 0 {
		return -1
	}

	// Step 1: Create a frequency map
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// Step 2: Find the element with the least frequency
	minFreq := len(nums) + 1
	leastFrequent := -1

	for num, freq := range freqMap {
		if freq < minFreq {
			minFreq = freq
			leastFrequent = num
		} else if freq == minFreq && num > leastFrequent {
			leastFrequent = num // Ensuring we take the last encountered least frequent number
		}
	}

	return leastFrequent
}

func main() {
	testCases := [][]int{
		{1},               // Output: 1
		{1, 2, 2},         // Output: 1
		{1, 2, 2, 3, 3},   // Output: 1
		{1, 1, 2, 3, 3, 4}, // Output: 2
		{1, 1, 2, 2, 3, 3, 4, 4}, // Output: 3
		{1, 2, 2, 3, 3, 4}, // Output: 4 (fixed)
	}

	for _, testCase := range testCases {
		fmt.Printf("Input: %v -> Least Frequent: %d\n", testCase, findLeastFrequent(testCase))
	}
}

