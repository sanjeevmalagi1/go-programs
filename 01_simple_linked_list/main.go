package main

import "fmt"

// Node represents a single element in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents the linked list with a head pointer
type LinkedList struct {
	head *Node
}

func (ll *LinkedList) insertAtBeginning(data int) {
	newNode := &Node{data: data, next: ll.head}

	ll.head = newNode
}

// insertAtEnd adds a new node at the end of the linked list
func (ll *LinkedList) insertAtEnd(data int) {
	newNode := &Node{data: data}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	last := ll.head
	for last.next != nil {
		last = last.next
	}

	last.next = newNode
}

func (ll *LinkedList) display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	ll := LinkedList{}

	// Insert nodes
	ll.insertAtBeginning(3)
	ll.insertAtBeginning(2)
	ll.insertAtBeginning(3)
	ll.insertAtBeginning(1)

	ll.insertAtEnd(4)

	ll.display()
}
