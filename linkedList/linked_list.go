package linkedList

type DoublyLinkedList struct {
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
func GetLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, nil, 0}
}

// Construct a linked list by an array of values and return the head
func GetLinkedListFromValues(vals []int) *DoublyLinkedList {
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
func (ll *DoublyLinkedList) Insert(val int) {
	newNode := &node{value: val}
	ll.insertNode(newNode)
}

// Insert at end by node
func (ll *DoublyLinkedList) insertNode(newNode *node) {
	ll.count++
	if ll.tail == nil {
		ll.head, ll.tail = newNode, newNode
	} else {
		ll.tail.next = newNode
		newNode.prev = ll.tail
		ll.tail = newNode
	}
}

// Insert at end by key
func (ll *DoublyLinkedList) getNode(index int) *node {
	node := ll.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

// Slice of the linked list
func (ll *DoublyLinkedList) Slice(begin int, end int) []int {
	slice := make([]int, end-begin+1)
	node := ll.head
	for i := begin; i <= end; i++ {
		slice[i] = node.value
		node = node.next
	}
	return slice
}

// Get the value in the linked list by index
func (ll *DoublyLinkedList) Get(index int) int {
	return ll.getNode(index).value
}

// Length of the linked list
func (ll *DoublyLinkedList) Len() int {
	return ll.count
}
