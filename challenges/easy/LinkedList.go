package easy

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insert(num int) {
	newNode := &Node{data: num, next: nil}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *LinkedList) print() {

	if list.head == nil {
		fmt.Println("No LinkedList found")
		return
	}

	current := list.head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}

	fmt.Printf("nil\n")

}

func LinkedListChallenge() {

	// Challenge 3: Linked List Implementation
	// Title: Challenge 3: Linked List Implementation
	// Description: Implement a singly linked list and insert elements.

	// Input:
	// linked_list := LinkedList{}
	// linked_list.insert(1)
	// linked_list.insert(2)
	// linked_list.insert(3)

	// Desired Output:
	// 1 -> 2 -> 3

	// Incorrect Output:
	// 1 -> 3 -> 2

	linked_list := LinkedList{}
	linked_list.insert(1)
	linked_list.insert(2)
	linked_list.insert(3)
	linked_list.insert(7)
	linked_list.insert(8)
	linked_list.insert(9)

	linked_list.print()

}
