package intermediate

import (
	"fmt"
)

type QNode struct {
	data int
	next *QNode
}

type Queue struct {
	head *QNode
	tail *QNode
}

func (q *Queue) enqueue(data int) {
	// Add to tail
	newNode := &QNode{data: data, next: nil}

	if q.head == nil && q.tail == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		last := q.tail
		q.tail = newNode
		last.next = q.tail
	}

}

func (q *Queue) dequeue() {
	// Remove from head, return value
	if q.head == nil && q.tail == nil {
		fmt.Println("Dequeue: No available Queue")
		return
	}
	first := q.head
	newHead := first.next
	q.head = newHead
	if newHead == nil {
		q.tail = nil
	}

	fmt.Println("Removed data: ", first.data)
}

func (q *Queue) print() {
	// Print nodes
	if q.head == nil {
		fmt.Println("Print: No Queue available")
		return
	}

	current := q.head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Printf("nil\n")

}

func (q *Queue) count() {
	count := 0
	current := q.head
	for current != nil {
		current = current.next
		count++
	}
	fmt.Println("Count: ", count)
}

func QueueChallenge() {

	// Challenge 5: Queue Implementation
	// Title: Challenge 5: Queue Implementation
	// Description: Implement a queue data structure and perform enqueue and dequeue operations.

	// Input:
	// queue := Queue{}
	// queue.enqueue(1)
	// queue.enqueue(2)
	// queue.dequeue()

	// Desired Output:
	// 1

	// Incorrect Output:
	// 2

	queue := Queue{}
	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)
	queue.enqueue(4)
	queue.dequeue()

	queue.count()

	queue.print()

}
