package linkedlist

import (
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	ml := Constructor()
	for i := 0; i < 5; i++ {
		ml.AddAtTail(i)
	}

	head := reverseBetween(ml.PreHead.Next, 2, 4)
	for head != nil {
		t.Log(head.Val)
		head = head.Next
	}
}
