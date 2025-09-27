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
