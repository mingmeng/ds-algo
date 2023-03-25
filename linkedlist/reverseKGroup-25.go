package linkedlist

func reverseKGroup(head *ListNode, k int) *ListNode {
	preHead := &ListNode{Next: head}

	prev, tail := preHead, preHead
	for prev.Next != nil {
		for i := 0; i < k; i++ {
			tail = tail.Next
			// 这里注意，一定是先往后移再判断
			// 如果是先判断再后移，一组k个的最后一个元素会判断不到。
			if tail == nil {
				return preHead.Next
			}
		}

		nex := tail.Next

		// 截断出一个临时链表
		tmpHead := prev.Next

		tail.Next = nil

		// 反转链表，此时新的 head 就是当前的 prev.Next
		prev.Next = reverseListIterate(tmpHead)
		// 此时临时链表的head是新tail
		tail = tmpHead

		// 将新 tail 的next 重新接上
		tail.Next = nex

		// prev 和 tail 放到同一个游标上面
		prev = tail

	}

	return preHead.Next
}
