package easy

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Insert(num int) {
	newNode := &Node{Data: num, Next: nil}

	if list.Head == nil {
		list.Head = newNode
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func (list *LinkedList) Print() {

	if list.Head == nil {
		fmt.Println("No LinkedList found")
		return
	}

	current := list.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Data)
		current = current.Next
	}

	fmt.Printf("nil\n")

}

func (list *LinkedList) Reverse() {
	if list.Head == nil {
		fmt.Println("No LinkedList found")
		return
	}

	var prev *Node = nil
	current := list.Head
	var next *Node = nil

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	list.Head = prev
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
	linked_list.Insert(1)
	linked_list.Insert(2)
	linked_list.Insert(3)
	linked_list.Insert(7)
	linked_list.Insert(8)
	linked_list.Insert(9)

	linked_list.Print()

}
