package binarytree

import (
	"container/list"
	"math"
)

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := list.New()
	cur := root
	var pre *TreeNode
	for stack.Len() > 0 || cur != nil {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			if pre != nil && cur.Val <= pre.Val {
				return false
			}
			pre = cur
			cur = cur.Right
		}
	}
	return true

}

var result int = math.MaxInt64
var pre *TreeNode

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	helper(root)

	return result
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}

	helper(root.Left)
	if pre != nil {
		result = min(result, root.Val-pre.Val)
	}
	pre = root
	helper(root.Right)
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
