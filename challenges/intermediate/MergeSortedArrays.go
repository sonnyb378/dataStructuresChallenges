package intermediate

import (
	"fmt"
)

func MergeSortedArrays(arr1 []int, arr2 []int) {

	// Challenge 8: Merge Sorted Arrays
	// Title: Challenge 8: Merge Sorted Arrays
	// Description: Merge two sorted arrays into a single sorted array.

	// Input:
	// arr1 := []int{1, 3, 5}
	// arr2 := []int{2, 4, 6}

	// Desired Output:
	// [1, 2, 3, 4, 5, 6]

	// Incorrect Output:
	// [1, 3, 2, 4, 5, 6]

	unsorted := append(arr1, arr2...)

	fmt.Println("Unsorted: ", unsorted)
	BubbleSort(unsorted)
	SelectionSort(unsorted)
	fmt.Println()
}

func BubbleSort(arr []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			next := i + 1
			if arr[i] > arr[next] {
				arr[i], arr[next] = arr[next], arr[i]
				swapped = true
			}
		}
	}
	fmt.Println()
	fmt.Println("Bubble Sort:")
	fmt.Println(arr)
}

func SelectionSort(arr []int) {
	marker := 0
	for marker < len(arr) {
		// fmt.Printf("%v \n", arr)
		for i := marker; i < len(arr); i++ {
			if arr[i] < arr[marker] {
				arr[i], arr[marker] = arr[marker], arr[i]
				// fmt.Printf("sub -> (%v) \n", arr)
			}
		}
		marker++

	}
	fmt.Println()
	fmt.Println("Selection Sort:")
	fmt.Println(arr)
}
