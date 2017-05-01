package linkedList

/*
2.3 Delete Middle Node: Delete a node in the middle (i.e., any node but the
first and last node, not necessarily the exact middle) of a singly linked list
given only access to that node.
e.g. Input: the node c from the linked list a->b->c->d->e->f
Result: nothing is returned, but the new linked list looks like a->b->d->e->f
*/
func (ll *DoubleLinkedList) removeNode(node *node) {
	if node == ll.head {
		ll.head = node.next
		ll.head.prev = nil
	} else if node.next == nil {
		ll.tail = node.prev
		ll.tail.next = nil
	} else {
		temp := ll.head
		for temp.next != node {
			temp = temp.next
		}

		temp.next = node.next
		node.next.prev = temp
	}

	node = nil
	ll.count--
}

func (ll *DoubleLinkedList) RemoveMiddleNode(index int) {
	ll.removeNode(ll.getNode(index))
}
