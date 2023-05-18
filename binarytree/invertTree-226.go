package binarytree

import "container/list"

// 翻转二叉树 递归
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}

	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp
	return root
}

// 翻转二叉树 迭代
func invertTreeIterate(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}

		node.Left, node.Right = node.Right, node.Left
	}

	return root
}
