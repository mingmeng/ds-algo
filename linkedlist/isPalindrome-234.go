package linkedlist

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	cur := head
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}

	mid := length / 2
	mid += length % 2
	fast := head
	for i := 0; i < mid; i++ {
		fast = fast.Next
	}

	now := head
	var pre *ListNode
	for i := 0; i < length/2; i++ {
		next := now.Next

		now.Next = pre
		pre = now
		now = next
	}

	for fast != nil {
		if fast.Val != pre.Val {
			return false
		}
		fast = fast.Next
		pre = pre.Next
	}

	return true
}
