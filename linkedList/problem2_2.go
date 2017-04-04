package linkedList

import "log"

/*
2.2. Return Kth to Last: Find the k th to last element of a singly linked list and return the value.
Without using prev could imitate singly linked list
*/
func (ll *DoubleLinkedList) KFromTail(k int) int {
	if k < 0 && k >= ll.count {
		log.Fatal("K is not within the length\n")
	}
	n := ll.head
	for i := 0; i < ll.count-k-1; i++ {
		n = n.next
	}
	return n.value
}
