package binarytree

func sumOfLeftLeavesRecur(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 判断当前节点是叶子结点的时候，其实也可以直接return 0 但是也可以不做
	// 这里暂时不做

	leftSum := sumOfLeftLeavesRecur(root.Left)
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			leftSum = root.Left.Val
		}
	}
	rightSum := sumOfLeftLeavesRecur(root.Right)

	return leftSum + rightSum
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	container := make([]*TreeNode, 0)
	container = append(container, root)

	var result int
	for len(container) > 0 {
		length := len(container)
		node := container[length-1]
		// 出栈一个
		container = container[:length-1]

		if node.Left != nil {
			if node.Left.Left == nil && node.Left.Right == nil {
				result += node.Left.Val
			}
		}

		if node.Right != nil {
			container = append(container, node.Right)
		}

		if node.Left != nil {
			container = append(container, node.Left)
		}

	}

	return result
}
