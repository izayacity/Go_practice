package linkedList

/*
2.7 Intersection: Given two (singly) linked lists, determine if the two lists
intersect. Return the intersecting node. Note that the intersection is defined
based on references, not value. That is, if the kth node of the first linked list
is the exact same node (by reference) as the jth node of the second linked list,
then they are intersecting.
*/
func AreIntersecting(l1, l2 *DoubleLinkedList) bool {
	n1, n2 := l1.head, l2.head

	if l1.count > l2.count {
		for i := 0; i < l1.count-l2.count; i++ {
			n1 = n1.next
		}
	} else if l1.count < l2.count {
		for i := 0; i < l2.count-l1.count; i++ {
			n2 = n2.next
		}
	}

	for n1 != nil {
		if n1 == n2 {
			return true
		}
		n1, n2 = n1.next, n2.next
	}

	return false
}
