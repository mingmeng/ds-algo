package binarytree

import "container/list"

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		lengthOfQueue := queue.Len()
		floor := make([]int, lengthOfQueue)
		for i := 0; i < lengthOfQueue; i++ {
			node, ok := queue.Remove(queue.Front()).(*TreeNode)
			if node != nil && ok {
				floor[i] = node.Val
			} else {
				floor[i] = -101
				continue
			}

			if node.Left != nil {
				queue.PushBack(node.Left)
			} else {
				queue.PushBack(nil)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			} else {
				queue.PushBack(nil)
			}

		}

		for i := 0; i < len(floor)/2; i++ {
			if floor[i] != floor[len(floor)-1-i] {
				return false
			}
		}
	}

	return true
}
