package linkedList

/*
2.8 Loop Detection: Given a circular linked list, implement an algorithm that
returns the node at the beginning of the loop.
DEFINITION:
Circular linked list: A (corrupt) linked list in which a node's next pointer
points to an earlier node, so as to make a loop in the linked list.
E.g. Input: A -> B -> C -> D -> E -> F -> D [the same C as earlier]
Output: C
*/
func (ll *DoubleLinkedList) FindLoopNode() *node {
	fast, slow := ll.head, ll.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.next == nil {
		return nil
	}

	slow = ll.head
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}
	return fast
}
