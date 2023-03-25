package linkedlist

import (
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	ml := Constructor()
	for i := 0; i < 5; i++ {
		ml.AddAtTail(i)
	}

	reverseKGroup(ml.PreHead.Next, 2)
}
