package LinkedList

import "Practice/utils"

type ListNode = utils.ListNode

func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	list := head
	rList := reverseList(list)

	for head != nil {
		if head.Val != rList.Val {
			return false
		}
		head = head.Next
		rList = rList.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	curr := &ListNode{Val: head.Val, Next: nil}
	list := head.Next
	for list != nil {
		node := &ListNode{Val: list.Val, Next: curr}
		curr = node
		list = list.Next
	}
	return curr
}
