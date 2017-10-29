package linkedList

import "fmt"
import "log"

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

// Convert linked list into an array
func (ll *DoubleLinkedList) Slice() []int {
	slice := make([]int, ll.count)
	node := ll.head
	for i := 0; i < ll.count; i++ {
		slice[i] = node.value
		node = node.next
	}
	return slice
}

// Get the value in the linked list by index
func (ll *DoubleLinkedList) Get(index int) int {
	return ll.getNode(index).value
}

// Get the value in the linked list by a node
func (ll *DoubleLinkedList) GetIndex(n *node) int {
	if n == nil {
		log.Fatal("Error: Nil node\n")
	}
	i := 0
	for temp := ll.head; temp != n; temp = temp.next {
		i++
	}
	return i
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
		for n := ll.head; n.next != nil; n = n.next {
			fmt.Print(n.value, " -> ")
		}
		fmt.Println(ll.tail.value)
	}
}

// Delete a node by value
func (ll *DoubleLinkedList) remove(value int) {
	for n := ll.head; n != nil; n = n.next {
		if n.value == value {
			ll.deleteNode(n)
		}
	}
}

// Delete a node by node
func (ll *DoubleLinkedList) deleteNode(n *node) {
	ll.count--
	if n == ll.head {
		ll.head = ll.head.next
		ll.head.prev = nil
	} else if n == ll.tail {
		ll.tail = ll.tail.prev
		ll.tail.next = nil
	} else {
		n.prev.next = n.next
		n.next.prev = n.prev
	}
	n = nil
}
