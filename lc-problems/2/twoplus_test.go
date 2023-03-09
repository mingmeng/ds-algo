package two

import "testing"

func TestTwo(t *testing.T) {
	l1 := &ListNode{Val: 0, Next: nil}
	l2 := &ListNode{Val: 0, Next: nil}
	addTwoNumbers(l1, l2)
}
