package linkedList

import "fmt"

type DoubleLinkedList struct {
	head  *node
	tail  *node
	count int
}

type node struct {
	value int
	next  *node
	prev  *node
}

// Construct an empty linked list and return the head
func GetLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{nil, nil, 0}
}

// Construct a linked list by an array of values and return the head
func GetLinkedListFromValues(vals []int) *DoubleLinkedList {
	ll := GetLinkedList()
	if len(vals) == 0 {
		return ll
	}
	for _, val := range vals {
		ll.Insert(val)
	}
	return ll
}

// Insert at end by value
func (ll *DoubleLinkedList) Insert(val int) {
	newNode := &node{value: val}
	ll.insertNode(newNode)
}

// Insert at end by node
func (ll *DoubleLinkedList) insertNode(newNode *node) {
	ll.count++
	if ll.tail == nil {
		ll.head, ll.tail = newNode, newNode
	} else {
		ll.tail.next = newNode
		newNode.prev = ll.tail
		ll.tail = newNode
	}
}

// Get node by key
func (ll *DoubleLinkedList) getNode(index int) *node {
	node := ll.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

// Slice of the linked list
func (ll *DoubleLinkedList) Slice(begin int, end int) []int {
	slice := make([]int, end-begin+1)
	node := ll.head
	for i := begin; i <= end; i++ {
		slice[i] = node.value
		node = node.next
	}
	return slice
}

// Get the value in the linked list by index
func (ll *DoubleLinkedList) Get(index int) int {
	return ll.getNode(index).value
}

// Length of the linked list
func (ll *DoubleLinkedList) Len() int {
	return ll.count
}

// print the linked list
func (ll *DoubleLinkedList) print() {
	if ll.count == 0 {
		fmt.Println("Empty linked list")
	} else {
		n := ll.head
		for n.next != nil {
			fmt.Print(n.value, " -> ")
			n = n.next
		}
		fmt.Println(n.value)
	}
}

// Delete a node by value
func (ll *DoubleLinkedList) deleteNode(value int) *DoubleLinkedList {
	n := ll.head

	if n.value == value {
		ll.head = ll.head.next
		ll.head.prev = nil
		return ll
	}

	for n.next != nil {
		if n.next.value == value {
			n.next = n.next.next
			n.next.next.prev = n.prev
			return ll
		}
		n = n.next
	}

	return ll
}
