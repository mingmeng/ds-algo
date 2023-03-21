package linkedlist

func removeElements(head *ListNode, val int) *ListNode {
	preHead := &ListNode{
		Val:  0,
		Next: head,
	}
	pos := preHead
	for pos.Next != nil {
		if pos.Next.Val == val {
			pos.Next = pos.Next.Next
		} else {
			pos = pos.Next
		}
	}
	return preHead.Next
}
