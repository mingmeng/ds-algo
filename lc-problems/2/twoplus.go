package two

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := new(ListNode)
	current := new(ListNode)

	head = current
	carry := 0
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		carry = sum / 10
		value := sum % 10

		//tmp := new(ListNode)
		//tmp.Val = value
		//tmp.Next = nil
		tmp := &ListNode{Val: value, Next: nil}

		current.Next = tmp
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{Val: 1, Next: nil}
	}

	return head.Next
}
