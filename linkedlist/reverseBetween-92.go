package linkedlist

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	preHead := &ListNode{Next: head}

	cur := preHead
	pos := 0

	var segStarts, _ *ListNode
	var segPrev, segNext *ListNode
	for cur != nil {
		if left == pos+1 {
			segPrev = cur
			segStarts = cur.Next
		}
		if right == pos {
			_ = cur
			segNext = cur.Next
			// 找到 right 之后就可以停了已经
			break
		}
		cur = cur.Next
		pos++
	}

	prev := segNext

	cur = segStarts
	for cur != segNext {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	segPrev.Next = prev

	return preHead.Next
}
