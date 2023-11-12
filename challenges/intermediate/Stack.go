package intermediate

import (
	"fmt"
)

type SNode struct {
	data int
	next *SNode
}

type Stack struct {
	pointer *SNode
}

func (st *Stack) push(data int) {
	// add on top of stack
	newNode := &SNode{data: data, next: nil}
	if st.pointer == nil {
		st.pointer = newNode
	} else {
		current := st.pointer
		st.pointer = &SNode{data: data, next: current}
	}
}

func (st *Stack) pop() {
	// remove on top of stack
	if st.is_empty() {
		fmt.Println("No Stack available")
		return
	}

	current := st.pointer
	newPointer := current.next
	for current != nil {
		current = current.next
	}
	st.pointer = newPointer
}

func (st *Stack) count() {
	if st.is_empty() {
		fmt.Println("No Stack available")
		return
	}
	count := 0
	current := st.pointer
	for current != nil {
		current = current.next
		count++
	}
	fmt.Println(count)
}

func (st *Stack) is_empty() bool {
	// check if stack is empty
	return st.pointer == nil
}
func (st *Stack) print() {
	// print stack
	if st.is_empty() {
		fmt.Println("No Stack available")
		return
	}
	current := st.pointer
	for current != nil {
		fmt.Printf(" %v \n", current.data)
		fmt.Println(" | ")
		fmt.Println(" v ")
		current = current.next
	}
	fmt.Println("None")

}

func StackChallenge() {

	// Challenge 4: Stack Implementation
	// Title: Challenge 4: Stack Implementation
	// Description: Create a stack data structure and perform push and pop operations.

	// Input:
	// stack := Stack{}
	// stack.push(1)
	// stack.push(2)
	// stack.pop()

	// Desired Output:
	// 1

	// Incorrect Output:
	// 2

	stack := Stack{}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(5)
	stack.pop()

	stack.print()
	stack.count()

}
