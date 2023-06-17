package binarytree

import "container/list"

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		length := stack.Len()

		floor := make([]int, 0)
		for i := 0; i < length; i++ {
			node := stack.Remove(stack.Front()).(*TreeNode)
			floor = append(floor, node.Val)

			if node.Left != nil {
				stack.PushBack(node.Left)
			}
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
		}
		result = append(result, floor)
	}

	for i := 0; i < len(result)/2; i++ {
		tmp := result[i]
		result[i] = result[len(result)-i-1]
		result[len(result)-i-1] = tmp
	}

	return result
}
