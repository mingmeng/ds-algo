package linked_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	preHead := &ListNode{
		Val:  0,
		Next: head,
	}
	fast, slow := head, preHead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return preHead.Next
}
