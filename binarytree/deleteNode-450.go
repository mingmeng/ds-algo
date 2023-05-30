package binarytree

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil || root.Val == key {
		return root
	}

	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left != nil && root.Right == nil {
			return root.Left
		} else if root.Right != nil && root.Left == nil {
			return root.Right
		} else {
			rightNode := root.Right
			for rightNode.Left != nil {
				rightNode = rightNode.Left
			}
			rightNode.Left = root.Left
			root = root.Right
			return root
		}
	}

	if root.Left != nil {
		root.Left = deleteNode(root.Left, key)
	}
	if root.Right != nil {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}
