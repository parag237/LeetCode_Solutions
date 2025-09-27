package utils

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Helper function to create a linked list from an array
func CreateLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	
	head := &ListNode{Val: values[0]}
	current := head
	
	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}
	
	return head
}

// Helper function to print linked list values
func PrintLinkedList(head *ListNode) {
	print("[")
	current := head
	for current != nil {
		print(current.Val)
		if current.Next != nil {
			print(", ")
		}
		current = current.Next
	}
	print("]")
}
