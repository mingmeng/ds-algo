package linkedlist

func swapPairs(head *ListNode) *ListNode {
	preHead := &ListNode{Val: 0, Next: head}
	cur := preHead

	for cur.Next != nil && cur.Next.Next != nil {
		prev := cur.Next
		tail := cur.Next.Next

		next := tail.Next

		cur.Next = tail
		tail.Next = prev
		prev.Next = next

		cur = cur.Next.Next
	}

	return preHead.Next
}
