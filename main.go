package main

import (
	"fmt"

	"github.com/sonnyb378/dataStructuresChallenges/challenges/intermediate"
)

func main() {

	// fmt.Println("Challenge 1: Array Sum")
	// arr := []int{1, 2, 3, 4, 5}
	// easy.ArraySum(arr)

	// fmt.Println("Challenge 2: Find Maximum Element")
	// arr := []int{
	// 	736, 195, 149, 970, 187, 467, 979, 534, 742, 192,
	// 	620, 941, 235, 848, 416, 243, 671, 972, 513, 763,
	// 	345, 635, 961, 374, 815, 50, 404, 529, 452, 104,
	// 	896, 841, 926, 538, 717, 708, 135, 157, 724, 743,
	// 	320, 867, 584, 358, 848, 16, 731, 645, 12, 69,
	// 	310, 23, 286, 912, 823, 99, 480, 719, 830, 181,
	// 	356, 515, 572, 476, 236, 558, 972, 763, 781, 38,
	// 	791, 739, 330, 248, 443, 915, 285, 341, 178, 483,
	// 	923, 890, 167, 744, 636, 558, 399, 448, 255, 532,
	// 	491, 985, 839, 526, 292, 219, 387, 536, 818, 659,
	// }
	// easy.FindMax(arr)

	// fmt.Println("Challenge 3: Linked List Implementation")
	// easy.LinkedListChallenge()

	// fmt.Println("Challenge 4: Stack Implementation")
	// easy.StackChallenge()

	// fmt.Println("Challenge 5: Queue Implementation")
	// easy.QueueChallenge()

	// fmt.Println("Challenge 6: Reverse Linked List Implementation")
	// intermediate.ReverseLinkedList()

	// fmt.Println("Challenge 7: Hash Map Implementation")
	// intermediate.HashMapChallenge()

	fmt.Println("Challenge 8: Merge Sorted Arrays")
	arr1 := []int{1, 4, 5, 6, 7, 12, 15, 19, 21, 56, 81, 101, 103}
	arr2 := []int{2, 3, 8, 9, 10, 45, 53, 69, 77, 82, 88, 90, 99}
	fmt.Println("arr1: ", arr1)
	fmt.Println("arr2: ", arr2)
	intermediate.MergeSortedArrays(arr1, arr2)
}
