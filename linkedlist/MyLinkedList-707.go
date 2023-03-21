package linkedlist

type MyLinkedList struct {
	PreHead *ListNode
	Size    int
}

func Constructor() MyLinkedList {
	return MyLinkedList{PreHead: &ListNode{
		Val:  0,
		Next: nil,
	}}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.Size || index < 0 {
		return -1
	}
	cur := this.PreHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}
	this.Size++
	cur := this.PreHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	newNode := &ListNode{Val: val, Next: cur.Next}
	cur.Next = newNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	cur := this.PreHead
	this.Size--
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
}
