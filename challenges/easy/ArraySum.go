package easy

import (
	"fmt"
)

func ArraySum(arr []int) {

	// Challenge 1: Array Sum
	// Title: Challenge 1: Array Sum
	// Description: Calculate the sum of elements in an array.

	// Input:
	// arr = [1, 2, 3, 4, 5]

	// Desired Output:
	// 15

	// Incorrect Output:
	// 10

	sum := 0
	for _, val := range arr {
		sum += val
	}

	fmt.Println("> ", sum)

}
