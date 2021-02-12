package linkedlist

// 递归方式链表反转
func reverseListRecur(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverseListRecur(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

// 迭代方式链表反转
func reverseListIterate(head *ListNode) *ListNode {
	var pre *ListNode
	now := head

	for now != nil {
		next := now.Next

		now.Next = pre
		pre = now
		now = next
	}

	return pre
}
