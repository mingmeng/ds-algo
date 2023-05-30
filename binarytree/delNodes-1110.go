package binarytree

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}

	var ans []*TreeNode
	var dfs func(root *TreeNode, to_delete []int) *TreeNode
	dfs = func(root *TreeNode, to_delete []int) *TreeNode {
		if root == nil {
			return nil
		}

		root.Left = dfs(root.Left, to_delete)
		root.Right = dfs(root.Right, to_delete)

		if !isDeleted(to_delete, root.Val) {
			return root
		}

		if root.Left != nil {
			ans = append(ans, root.Left)
		}
		if root.Right != nil {
			ans = append(ans, root.Right)
		}
		return nil
	}
	if dfs(root, to_delete) != nil {
		ans = append(ans, root)
	}

	return ans
}

func isDeleted(to_delete []int, val int) bool {
	for _, v := range to_delete {
		if v == val {
			return true
		}
	}
	return false
}
