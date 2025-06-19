/*
You are given a list of 11 integers. Your task is to sort only a specific portion of the list—namely,
the elements from index 2 to index 8 (inclusive)—while keeping the rest of the list unchanged.

// take 10 random number
// 1 [2 - 8] ..10 sort
// print form index 2 to 8 with sort other keep index
// 1 [sorted list] 9 10
// input : 6,2,1,5,6,7,9,0,12,50,4
// output: 6 [0,1,2,5,6,7,9,12] 50,4
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{6, 2, 1, 5, 6, 7, 9, 0, 12, 50, 4}

	startPoint := 2
	endPoint := 8

	sortedPart := append([]int{}, array[startPoint:endPoint+1]...)

	sort.Ints(sortedPart)

	copy(array[startPoint:endPoint+1], sortedPart)

	fmt.Println("--> ", array)
}
