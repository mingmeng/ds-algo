package binarytree

import "container/list"

func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)

	stack := list.New()

	maxCount := 0
	count := 0
	var pre *TreeNode

	cur := root
	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			if pre == nil {
				count = 1
			} else if cur.Val == pre.Val {
				count++
			} else {
				count = 1
			}

			if count == maxCount {
				result = append(result, cur.Val)
			}
			if count > maxCount {
				maxCount = count
				result = make([]int, 0)
				result = append(result, cur.Val)
			}
			pre = cur
			cur = cur.Right
		}
	}
	return result
}
