package main

import (
	"fmt"
)

func allTwoSum(nums []int, target int) [][2]int {
	pairs := [][2]int{}
	seen := make(map[int]bool)
	added := make(map[string]bool) // To avoid duplicates

	for _, num := range nums {
		complement := target - num
		if seen[complement] {
			// Ensure uniqueness of the pair regardless of order
			a, b := num, complement
			if a > b {
				a, b = b, a
			}
			key := fmt.Sprintf("%d:%d", a, b)
			if !added[key] {
				pairs = append(pairs, [2]int{a, b})
				added[key] = true
			}
		}
		seen[num] = true
	}

	return pairs
}

func main() {
	arr := []int{1, 2, 3, 2, 4, 5, 6, 3, 7}
	target := 6

	pairs := allTwoSum(arr, target)
	if len(pairs) == 0 {
		fmt.Println("No matching pairs found.")
	} else {
		fmt.Println("Matching pairs:")
		for _, pair := range pairs {
			fmt.Printf("%d + %d = %d\n", pair[0], pair[1], target)
		}
	}
}
