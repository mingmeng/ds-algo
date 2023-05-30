package binarytree

import "container/list"

func flattenIterate(root *TreeNode) {
	if root == nil || root.Left == nil && root.Right == nil {
		return
	}

	stack := list.New()
	stack.PushBack(root)

	var flattenPre *TreeNode

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)

		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}

		if flattenPre != nil {
			flattenPre.Right = node
			flattenPre.Left = nil
			flattenPre = node
		}
	}
}

func flattenRecur(root *TreeNode) {
	if root == nil {
		return
	}

	flattenRecur(root.Left)
	flattenRecur(root.Right)

	newRight := root.Right
	root.Left, root.Right = nil, root.Left
	for root.Right != nil {
		root = root.Right
	}

	root.Right = newRight

}
