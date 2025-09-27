package LinkedList

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := &ListNode{Val: head.Val, Next: nil}
	list := head.Next
	for list != nil {
		node := &ListNode{Val: list.Val, Next: curr}
		curr = node
		list = list.Next
	}
	return curr
}
