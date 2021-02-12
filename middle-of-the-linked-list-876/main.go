package middle_of_the_linked_list_876

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	fast:=head
	slow:=head
	for fast!=nil&&fast.Next!=nil {
		slow=slow.Next
		fast=fast.Next.Next
	}
	return slow
}
