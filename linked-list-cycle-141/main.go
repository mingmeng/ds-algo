package linked_list_cycle_141

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针原理
// https://leetcode-cn.com/problems/linked-list-cycle/solution/xiang-jie-wei-shi-yao-yong-yi-bu-liang-b-i6xo/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fastP := head
	slowP := head

	for fastP != nil && fastP.Next != nil {
		slowP = slowP.Next
		fastP = fastP.Next.Next

		if slowP == fastP {
			return true
		}
	}
	return false
}
