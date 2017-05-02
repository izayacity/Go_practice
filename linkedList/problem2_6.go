package linkedList

/*
2.6 Palindrome: Implement a function to check if a linked list is a palindrome.
*/
func (ll *DoubleLinkedList) IsPalindrome() bool {
	if ll.head == nil || ll.head == ll.tail {
		return true
	}

	node1 := ll.head
	node2 := ll.tail

	for node1 != node2 && node1.next != node2 {
		if node1.value != node2.value {
			return false
		}

		node1 = node1.next
		node2 = node2.prev
	}
	return true
}
