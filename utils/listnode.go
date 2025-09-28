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

// Helper function to create a circular linked list from an array
// pos is the index to which the tail of the list will point to form the circular list
// pos = -1 means no cycle (regular linked list)
func CreateCircularLinkedList(values []int, pos int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	// Create regular linked list first
	head := &ListNode{Val: values[0]}
	current := head
	var nodes []*ListNode
	nodes = append(nodes, head)

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
		nodes = append(nodes, current)
	}

	// Create cycle if pos is valid
	if pos >= 0 && pos < len(values) {
		current.Next = nodes[pos]
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

// Helper function to print circular linked list values (with cycle detection)
func PrintCircularLinkedList(head *ListNode, maxNodes int) {
	if head == nil {
		print("[]")
		return
	}

	print("[")
	current := head
	count := 0
	for current != nil && count < maxNodes {
		print(current.Val)
		if current.Next != nil && count < maxNodes-1 {
			print(", ")
		}
		current = current.Next
		count++
	}
	if count >= maxNodes {
		print(", ... (cycle detected)")
	}
	print("]")
}

func IsDuplicateList(head1 *ListNode, head2 *ListNode) bool {
	for head1 != nil && head2 != nil {
		if head1.Val != head2.Val {
			return false
		}
		head1 = head1.Next
		head2 = head2.Next
	}
	if head1 == nil && head2 == nil {
		return true
	}
	return false
}

// Helper function to detect if a linked list has a cycle
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

// Helper function to find the cycle length in a circular linked list
func GetCycleLength(head *ListNode) int {
	if !HasCycle(head) {
		return 0
	}

	// Find the meeting point
	slow := head
	fast := head
	for {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	// Count the cycle length
	length := 1
	current := slow.Next
	for current != slow {
		length++
		current = current.Next
	}

	return length
}
