package intermediate

import "github.com/sonnyb378/dataStructuresChallenges/challenges/easy"

func ReverseLinkedList() {
	// Challenge 6: Reverse a Linked List
	// Title: Challenge 6: Reverse a Linked List
	// Description: Reverse a singly linked list.

	// Input:
	// linked_list = LinkedList()
	// linked_list.insert(1)
	// linked_list.insert(2)
	// linked_list.insert(3)

	// Desired Output:
	// 3 -> 2 -> 1

	// Incorrect Output:
	// 1 -> 2 -> 3

	linked_list := easy.LinkedList{}
	linked_list.Insert(1)
	linked_list.Insert(2)
	linked_list.Insert(3)
	linked_list.Insert(4)
	linked_list.Insert(5)
	linked_list.Print()

	linked_list.Reverse()
	linked_list.Print()

}
