package easy

import (
	"fmt"
)

func FindMax(arr []int) {

	// Challenge 2: Find Maximum Element
	// Title: Challenge 2: Find Maximum Element
	// Description: Find the maximum element in an array.

	// Input:
	// arr = [8, 2, 10, 4, 5]

	// Desired Output:
	// 10

	// Incorrect Output:
	// 8

	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	fmt.Println("Max: ", max)

}
