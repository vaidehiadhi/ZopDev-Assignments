package Solutions

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) InsertLast(data int) {
	newNode := &Node{data: data}
	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) DeleteLast() {
	if ll.head == nil {
		fmt.Println("List is empty")
		return
	}

	if ll.head.next == nil {
		ll.head = nil
		return
	}

	current := ll.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
}

func (ll *LinkedList) Display() {
	for current := ll.head; current != nil; current = current.next {
		fmt.Printf("%d -> ", current.data)
	}
	fmt.Println("nil")
}
