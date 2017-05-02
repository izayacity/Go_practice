package linkedList

/*
2.5 Sum Lists: You have two numbers represented by a linked list, where each
node contains a single digit. The digits are stored in reverse order, such that
the 1's digit is at the head of the list. Write a function that adds the two
numbers and returns the sum as a linked list.
E.g. Input: (7 -> 1 -> 6) + (5 -> 9 -> 2). That is, 617 + 295.
Output: 2 -> 1 -> 9. that is, 912
Follow up
Suppose the digits are stored in forward order. Repeat the above problem.
E.g. Input: (6 -> 1 -> 7) + (2 -> 9 -> 5). That is, 617 + 295.
Output: 9 -> 1 -> 2. That is, 912.
*/

func AddTwoLists(l1, l2 *DoubleLinkedList) int {
	value1 := ReverseInt(ListToInt(l1))
	value2 := ReverseInt(ListToInt(l2))
	return value1 + value2
}

func ListToInt(l_list *DoubleLinkedList) int {
	if l_list.head == nil {
		return 0
	}

	temp := l_list.head
	result := temp.value

	for temp.next != nil {
		temp = temp.next
		result = result*10 + temp.value
	}
	return result
}

func ReverseInt(number int) int {
	result := 0

	for number != 0 {
		result = result*10 + number%10
		number /= 10
	}
	return result
}
