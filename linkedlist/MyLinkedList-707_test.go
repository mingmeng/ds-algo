package linkedlist

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	t.Log(obj.Get(1))
	obj.DeleteAtIndex(1)
	t.Log(obj.Get(1))
	t.Log("-------------------------queue--------------")
	for cur := obj.PreHead.Next; cur != nil; cur = cur.Next {
		t.Log(cur.Val)
	}
}
