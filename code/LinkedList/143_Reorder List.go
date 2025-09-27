package LinkedList

func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Find the middle of the list
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Split the list into two halves
	secondList := slow.Next
	slow.Next = nil

	// Reverse the second half
	secondList = reverseList(secondList)

	// Merge the two halves alternately
	first := head
	second := secondList
	for second != nil {
		temp1 := first.Next
		temp2 := second.Next

		first.Next = second
		second.Next = temp1

		first = temp1
		second = temp2
	}
}
