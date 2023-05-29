package binarytree

func hasPathSumHelper(root *TreeNode, count int) bool {
	// 如果已经到达叶子节点，那么直接判断当前count 是否等于0
	if root.Left == nil && root.Right == nil {
		if count-root.Val == 0 {
			return true
		} else {
			// 返回false
			return false
		}
	}
	// 如果不是叶子节点，处理左边和右边
	if root.Left != nil {
		count -= root.Val
		if hasPathSumHelper(root.Left, count) {
			return true
		}
		count += root.Val
	}
	if root.Right != nil {
		count -= root.Val
		if hasPathSumHelper(root.Right, count) {
			return true
		}
		count += root.Val
	}

	return false
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return hasPathSumHelper(root, targetSum)
}
